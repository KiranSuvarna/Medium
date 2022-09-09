package main

import (
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

var websites = []string{
	"https://twitter.com",
	"https://www.google.co.in",
	"https://medium.com",
}

func pingWebsite(url string) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("%s looks down. Reason: %v", url, err)
	}

	log.Printf("%s is up. Status: %s", url, resp.Status)
}

func main() {
	for _, website := range websites {
		go pingWebsite(website)
		wg.Add(1)
	}
	wg.Wait()
}
