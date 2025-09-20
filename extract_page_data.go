package main

import "net/url"

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	h1 := getH1FromHTML(html)
	firstParagraph := getFirstParagraphFromHTML(html)
	parsedURL, err := url.Parse(pageURL)

	if err != nil {
		return PageData{
			URL:            pageURL,
			H1:             h1,
			FirstParagraph: firstParagraph,
			OutgoingLinks:  nil,
			ImageURLs:      nil,
		}
	}

	links, err := getURLsFromHTML(html, parsedURL)
	if err != nil {
		links = []string{}
	}

	images, err := getImagesFromHTML(html, parsedURL)
	if err != nil {
		images = []string{}
	}

	return PageData{
		URL:            pageURL,
		H1:             h1,
		FirstParagraph: firstParagraph,
		OutgoingLinks:  links,
		ImageURLs:      images,
	}
}
