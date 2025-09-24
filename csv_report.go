package main

import (
	"encoding/csv"
	"os"
	"sort"
	"strings"
)

func writeCSVReport(pages map[string]PageData, filename string) error {
	f, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	if err := w.Write([]string{
		"page_url",
		"h1",
		"first_paragraph",
		"outgoing_link_urls",
		"image_urls",
	}); err != nil {
		return err
	}

	keys := make([]string, 0, len(pages))

	for k := range pages {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, url := range keys {
		page := pages[url]

		outgoing := strings.Join(page.OutgoingLinks, ";")
		images := strings.Join(page.ImageURLs, ";")

		record := []string{
			page.URL,
			page.H1,
			page.FirstParagraph,
			outgoing,
			images,
		}

		if err := w.Write(record); err != nil {
			return err
		}
	}

	w.Flush()

	return w.Error()
}
