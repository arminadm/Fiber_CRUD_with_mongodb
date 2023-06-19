package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestCSVToMongo(t *testing.T) {
	// Send the POST request to the endpoint
	endpoint := "/csv-to-mongo"
	resp, err := http.Post(fmt.Sprintf("http://localhost:3000%s", endpoint), "application/json", nil)
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			t.Fatalf("In order to use this test, Server must be up and running, please make sure that backend and database container is up and running")
		}
		t.Fatalf("Failed to send request: %s", err.Error())
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	} else {
		log.Printf("response status code: %d\n", resp.StatusCode)
	}
}
