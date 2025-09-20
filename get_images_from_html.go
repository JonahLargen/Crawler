package main

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))

	if err != nil {
		return nil, err
	}

	images := make([]string, 0)

	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, exists := s.Attr("src")

		if !exists || strings.TrimSpace(src) == "" {
			return
		}

		src = strings.TrimSpace(src)
		parsed, err := url.Parse(src)

		if err != nil {
			return
		}

		absURL := parsed

		if !parsed.IsAbs() && baseURL != nil {
			absURL = baseURL.ResolveReference(parsed)
		}

		images = append(images, absURL.String())
	})

	return images, nil
}
