package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func sendRequest(id int, requestId int, testUrl string) int {
	log.Println("Worker", id, "is launching request", requestId)
	_, err := http.Get(testUrl)
	log.Println("Worker", id, "finished request", requestId)
	if err != nil {
		return 0
	}
	return 1
}

func main() {
	numberRequestsPtr := flag.Int("n", 1, "Number of requests to make")
	numberConcurrentConnectionsPtr := flag.Int("c", 1, "Number of concurrent connections")
	flag.Parse()

	testUrl := os.Args[len(os.Args)-1]

	log.Println("Url to test =", testUrl)
	log.Println("Number of requests =", *numberRequestsPtr)
	log.Println("Concurrent requests =", *numberConcurrentConnectionsPtr)

	requests := make(chan int)
	results := make(chan int, *numberRequestsPtr)

	for connectionId := 0; connectionId < *numberConcurrentConnectionsPtr; connectionId++ {
		go func(connectionId int) {
			for requestId := range requests {
				results <- sendRequest(connectionId, requestId, testUrl)
			}
		}(connectionId)
	}

	for requestId := 0; requestId < *numberRequestsPtr; requestId++ {
		requests <- requestId
	}
	close(requests)

	for resultPosition := 0; resultPosition < *numberRequestsPtr; resultPosition++ {
		<-results
	}
}
