package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))

	if err != nil {
		return nil, err
	}

	urls := make([]string, 0)

	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")

		if !exists || strings.TrimSpace(href) == "" {
			return
		}

		href = strings.TrimSpace(href)
		parsed, err := url.Parse(href)

		if err != nil {
			return
		}

		absURL := parsed

		if !parsed.IsAbs() && baseURL != nil {
			absURL = baseURL.ResolveReference(parsed)
		}

		urls = append(urls, absURL.String())
	})

	return urls, nil
}
