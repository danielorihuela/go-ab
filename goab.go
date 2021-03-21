package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/danielorihuela/goab/logger"
)

var log = logger.New(false, logger.DebugLevel)

func sendRequest(client *http.Client, testUrl string) int {
	resp, err := client.Get(testUrl)
	if err != nil {
		log.Error(err)
		return 0
	}
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	return 1
}

func main() {
	numberRequestsPtr := flag.Int("n", 1, "Number of requests to make")
	numberConcurrentConnectionsPtr := flag.Int("c", 1, "Number of concurrent connections")
	keepAlivePtr := flag.Bool("k", false, "Activate keep alive HTTP feature")
	flag.Parse()

	testUrl := os.Args[len(os.Args)-1]

	log.Debug("Url to test =", testUrl)
	log.Debug("Number of requests =", *numberRequestsPtr)
	log.Debug("Concurrent requests =", *numberConcurrentConnectionsPtr)
	log.Debug("Keep Alive HTTP is activated =", *keepAlivePtr)

	requests := make(chan int)
	results := make(chan int, *numberRequestsPtr)

	transport := &http.Transport{DisableKeepAlives: !*keepAlivePtr}
	client := &http.Client{Transport: transport}

	startRequests := time.Now()
	for connectionId := 0; connectionId < *numberConcurrentConnectionsPtr; connectionId++ {
		go func(connectionId int) {
			for range requests {
				results <- sendRequest(client, testUrl)
			}
		}(connectionId)
	}

	for requestId := 0; requestId < *numberRequestsPtr; requestId++ {
		requests <- requestId
	}
	close(requests)

	totalResults := 0
	for resultPosition := 0; resultPosition < *numberRequestsPtr; resultPosition++ {
		totalResults += <-results
	}
	timeTaken := time.Since(startRequests).Seconds()

	fmt.Println("Total time of program in seconds", timeTaken)
	fmt.Println("Requests per second", float64(*numberRequestsPtr)/timeTaken)
	fmt.Println("Time per request (mean)", float64(*numberConcurrentConnectionsPtr)*timeTaken*1000/float64(*numberRequestsPtr))
	fmt.Println("Time per request (mean, across all concurrent requests)", timeTaken*1000/float64(*numberRequestsPtr))

	fmt.Println("Errored responses =", *numberRequestsPtr-totalResults)
	errorPercentage := (1 - (float64(totalResults) / float64(*numberRequestsPtr))) * 100
	fmt.Println("Errored responses percentage =", strconv.FormatFloat(errorPercentage, 'f', 2, 64)+"%")
}
