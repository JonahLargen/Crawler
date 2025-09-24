package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]PageData
	seen               map[string]struct{}
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func (cfg *config) crawlPage(rawCurrentURL string) {
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

	cfg.mu.Lock()

	if _, exists := cfg.seen[normalizedURL]; exists {
		cfg.mu.Unlock()
		return
	}

	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}

	cfg.seen[normalizedURL] = struct{}{}
	cfg.mu.Unlock()

	cfg.wg.Add(1)

	go func(rawURL, normURL string) {
		cfg.concurrencyControl <- struct{}{}

		defer func() {
			<-cfg.concurrencyControl
			cfg.wg.Done()
		}()

		fmt.Printf("crawling %s\n", rawURL)

		htmlBody, err := getHTML(rawURL)

		if err != nil {
			fmt.Printf("Error - getHTML: %v\n", err)
			return
		}

		pageData := extractPageData(htmlBody, rawURL)

		cfg.mu.Lock()

		if len(cfg.pages) < cfg.maxPages {
			cfg.pages[normURL] = pageData
		}

		cfg.mu.Unlock()

		for _, nextURL := range pageData.OutgoingLinks {
			cfg.crawlPage(nextURL)
		}
	}(rawCurrentURL, normalizedURL)
}
