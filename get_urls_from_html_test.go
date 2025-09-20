package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	base, _ := url.Parse("https://example.com/dir/")
	tests := []struct {
		name     string
		html     string
		base     *url.URL
		expected []string
	}{
		{
			name: "absolute and relative links",
			html: `<a href="https://example.com/abs">Abs</a>
				<a href="/rel">Rel</a>
				<a href="foo/bar">Nested</a>`,
			base: base,
			expected: []string{
				"https://example.com/abs",
				"https://example.com/rel",
				"https://example.com/dir/foo/bar",
			},
		},
		{
			name: "ignore empty and invalid href",
			html: `<a href="">Empty</a>
				<a>No href</a>
				<a href="http://valid.com">Valid</a>
				<a href=" :badurl">Bad</a>`,
			base: base,
			expected: []string{
				"http://valid.com",
			},
		},
		{
			name:     "no links",
			html:     `<div>No links here!</div>`,
			base:     base,
			expected: []string{},
		},
		{
			name: "relative link with ..",
			html: `<a href="../up">Up Dir</a>`,
			base: base,
			expected: []string{
				"https://example.com/up",
			},
		},
		{
			name: "fragment and query",
			html: `<a href="/foo?bar=baz#frag">Link</a>`,
			base: base,
			expected: []string{
				"https://example.com/foo?bar=baz#frag",
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.html, tc.base)
			if err != nil {
				t.Errorf("Test %v - %s FAIL: unexpected error: %v", i, tc.name, err)
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected %v, got %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
