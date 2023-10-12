package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	url := "https://von327fk7jq6f0ls3n5s824m4da4ywml.burp.schellman.info" // Replace with the URL you want to reach

	// Send an HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}

	// Print the response body
	fmt.Printf("Response from %s:\n%s\n", url, string(body))
}
