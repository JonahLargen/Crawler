package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("Usage: ./crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}

	rawBaseURL := args[0]

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil || maxConcurrency < 1 {
		fmt.Println("maxConcurrency must be a positive integer")
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil || maxPages < 1 {
		fmt.Println("maxPages must be a positive integer")
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s (max concurrency: %d, max pages: %d)...\n",
		rawBaseURL, maxConcurrency, maxPages)

	parsedBaseURL, _ := url.Parse(rawBaseURL)

	cfg := &config{
		pages:              make(map[string]int),
		baseURL:            parsedBaseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
