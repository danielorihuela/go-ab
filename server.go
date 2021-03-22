package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func launchServer() {
	http.HandleFunc("/", hello)

	http.ListenAndServe(":1234", nil)
}
