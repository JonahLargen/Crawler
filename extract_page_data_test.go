package main

import (
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		pageURL  string
		expected PageData
	}{
		{
			name: "basic page with everything",
			html: `
				<html>
					<head><title>Test</title></head>
					<body>
						<h1>Test Heading</h1>
						<main>
							<p>This is the main paragraph.</p>
						</main>
						<a href="https://external.com/abc">External</a>
						<a href="/relative/link">Relative Link</a>
						<img src="https://img.com/image.jpg">
						<img src="/img2.png">
					</body>
				</html>
			`,
			pageURL: "https://test.com/section/",
			expected: PageData{
				URL:            "https://test.com/section/",
				H1:             "Test Heading",
				FirstParagraph: "This is the main paragraph.",
				OutgoingLinks: []string{
					"https://external.com/abc",
					"https://test.com/relative/link",
				},
				ImageURLs: []string{
					"https://img.com/image.jpg",
					"https://test.com/img2.png",
				},
			},
		},
		{
			name: "no h1, no main, relative links and images",
			html: `
				<html>
					<body>
						<p>First para not in main.</p>
						<a href="foo">Foo Link</a>
						<img src="bar.jpg">
					</body>
				</html>
			`,
			pageURL: "https://abc.com/path/",
			expected: PageData{
				URL:            "https://abc.com/path/",
				H1:             "",
				FirstParagraph: "First para not in main.",
				OutgoingLinks:  []string{"https://abc.com/path/foo"},
				ImageURLs:      []string{"https://abc.com/path/bar.jpg"},
			},
		},
		{
			name:    "empty html",
			html:    ``,
			pageURL: "https://empty.com/",
			expected: PageData{
				URL:            "https://empty.com/",
				H1:             "",
				FirstParagraph: "",
				OutgoingLinks:  []string{},
				ImageURLs:      []string{},
			},
		},
		{
			name:    "broken html and malformed URL",
			html:    `<h1>Oops`,
			pageURL: ":badurl",
			expected: PageData{
				URL:            ":badurl",
				H1:             "Oops",
				FirstParagraph: "",
				OutgoingLinks:  nil,
				ImageURLs:      nil,
			},
		},
		{
			name: "multiple links/images, skip missing attributes",
			html: `
				<a href="">Empty</a>
				<a>No href</a>
				<a href="http://foo.com">Foo</a>
				<img>No src</img>
				<img src="">
				<img src="img.png">
			`,
			pageURL: "https://site.com/",
			expected: PageData{
				URL:            "https://site.com/",
				H1:             "",
				FirstParagraph: "",
				OutgoingLinks:  []string{"http://foo.com"},
				ImageURLs:      []string{"https://site.com/img.png"},
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := extractPageData(tc.html, tc.pageURL)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %d - %s FAIL:\nExpected: %#v\nActual:   %#v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
