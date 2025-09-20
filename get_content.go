package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1FromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil {
		return ""
	}

	h1 := doc.Find("h1").First()

	if h1.Length() == 0 {
		return ""
	}

	return strings.TrimSpace(h1.Text())
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil {
		return ""
	}

	mainSel := doc.Find("main").First()

	var pSel *goquery.Selection

	if mainSel.Length() > 0 {
		pSel = mainSel.Find("p").First()

		if pSel.Length() == 0 {
			pSel = doc.Find("p").First()
		}
	} else {
		pSel = doc.Find("p").First()
	}

	if pSel.Length() == 0 {
		return ""
	}

	return strings.TrimSpace(pSel.Text())
}
