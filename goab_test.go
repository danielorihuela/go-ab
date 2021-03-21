package main

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestPageIsNotReachableStatusCode404(t *testing.T) {
	result := pageIsNotReachable(&http.Response{StatusCode: 404}, nil)

	expected := true
	if result != expected {
		t.Errorf("Result was %t, but we expected %t", result, expected)
	}
}

func TestPageIsNotReachableError(t *testing.T) {
	err := errors.New("test error")
	result := pageIsNotReachable(nil, err)

	expected := true
	if result != expected {
		t.Errorf("Result was %t, but we expected %t", result, expected)
	}
}

func TestPageIsNotReachableWhenThereIsNoProblem(t *testing.T) {
	result := pageIsNotReachable(&http.Response{StatusCode: 200}, nil)

	expected := false
	if result != expected {
		t.Errorf("Result was %t, but we expected %t", result, expected)
	}
}

func TestConcurrentConnectionsWillNotHaveRequest(t *testing.T) {
	var tests = []struct {
		concurrency, requests int
		expectedResult        bool
	}{
		{10, 2, true},
		{2, 10, false},
		{10, 10, false},
	}

	for _, testData := range tests {
		testname := fmt.Sprintf("Concurrent connections %d and number of requests %d", testData.concurrency, testData.requests)
		t.Run(testname, func(t *testing.T) {
			result := concurrentConnectionsWillNotHaveRequests(testData.concurrency, testData.requests)

			if result != testData.expectedResult {
				t.Errorf("Result was %t, but we expected %t", result, testData.expectedResult)
			}
		})
	}
}
