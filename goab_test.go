package main

import (
	"errors"
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
