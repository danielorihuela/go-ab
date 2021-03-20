package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func sendRequest(requestId int, testUrl string) {
	log.Println("Launching request", requestId)
	_, err := http.Get(testUrl)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("Finished")
}

func main() {
	numberRequestsPtr := flag.Int("n", 1, "Number of requests to make")
	flag.Parse()

	testUrl := os.Args[len(os.Args)-1]

	log.Println("Url to test =", testUrl)
	log.Println("Number of requests =", *numberRequestsPtr)

	for requestId := 1; requestId <= *numberRequestsPtr; requestId++ {
		sendRequest(requestId, testUrl)
	}
}
