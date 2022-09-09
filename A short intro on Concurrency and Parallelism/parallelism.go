package main

import (
	"log"
	"net/http"
	"runtime"
	"runtime/pprof"
	"sync"
)

var wg sync.WaitGroup
var threadProfile = pprof.Lookup("threadcreate")

var websites = []string{
	"https://twitter.com",
	"https://www.google.co.in",
	"https://medium.com",
}

func pingWebsite(url string) {
	runtime.LockOSThread()
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("%s looks down. Reason: %v", url, err)
	}

	log.Printf("%s is up. Status: %s", url, resp.Status)
	runtime.UnlockOSThread()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("Number of threads before: ", threadProfile.Count())

	for _, website := range websites {
		wg.Add(1)
		go pingWebsite(website)
	}

	wg.Wait()
	log.Println("Number of threads after: ", threadProfile.Count())
}
