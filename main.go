package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	parsedBaseURL, _ := url.Parse(rawBaseURL)

	cfg := &config{
		pages:              make(map[string]int),
		baseURL:            parsedBaseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 5), // max 1 concurrent requests
		wg:                 &sync.WaitGroup{},
	}

	cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
