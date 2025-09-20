package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func (cfg *config) addPageVisit(normalizedURL string) bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, exists := cfg.pages[normalizedURL]; exists {
		cfg.pages[normalizedURL]++
		return false
	}
	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	cfg.wg.Add(1)
	go func() {
		cfg.concurrencyControl <- struct{}{}
		defer func() {
			<-cfg.concurrencyControl
			cfg.wg.Done()
		}()

		currentURL, err := url.Parse(rawCurrentURL)
		if err != nil {
			fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
			return
		}

		if currentURL.Hostname() != cfg.baseURL.Hostname() {
			return
		}

		normalizedURL, err := normalizeURL(rawCurrentURL)
		if err != nil {
			fmt.Printf("Error - normalizeURL: %v\n", err)
			return
		}

		if !cfg.addPageVisit(normalizedURL) {
			return
		}

		cfg.mu.Lock()
		if len(cfg.pages) > cfg.maxPages {
			cfg.mu.Unlock()
			return
		}
		cfg.mu.Unlock()

		fmt.Printf("crawling %s\n", rawCurrentURL)

		htmlBody, err := getHTML(rawCurrentURL)
		if err != nil {
			fmt.Printf("Error - getHTML: %v\n", err)
			return
		}

		nextURLs, err := getURLsFromHTML(htmlBody, cfg.baseURL)
		if err != nil {
			fmt.Printf("Error - getURLsFromHTML: %v\n", err)
			return
		}

		for _, nextURL := range nextURLs {
			cfg.crawlPage(nextURL)
		}
	}()
}
