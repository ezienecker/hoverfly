package main

import (
	"net/http"
	"testing"
)

func TestReconstructRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)

	// changing payload so we don't have to call middleware
	request := requestDetails{
		Path:        "/random-path",
		Method:      "POST",
		Query:       "?foo=bar",
		Destination: "changed.destination.com",
	}
	payload := Payload{Request: request}

	c := NewConstructor(req, payload)
	newRequest := c.reconstructRequest()
	expect(t, newRequest.Method, "POST")
	expect(t, newRequest.URL.Path, "/random-path")
	expect(t, newRequest.Host, "changed.destination.com")
	expect(t, newRequest.URL.RawQuery, "?foo=bar")
}

func TestReconstructRequestEmptyPayload(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)

	payload := Payload{}
	c := NewConstructor(req, payload)
	newRequest := c.reconstructRequest()
	expect(t, newRequest.Method, "")
	expect(t, newRequest.Host, "")

}
