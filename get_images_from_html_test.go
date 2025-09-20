package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetImagesFromHTML(t *testing.T) {
	base, _ := url.Parse("https://img.com/path/")
	tests := []struct {
		name     string
		html     string
		base     *url.URL
		expected []string
	}{
		{
			name: "absolute and relative src",
			html: `<img src="https://img.com/abs.jpg">
				<img src="/rel.png">
				<img src="foo/bar.gif">`,
			base: base,
			expected: []string{
				"https://img.com/abs.jpg",
				"https://img.com/rel.png",
				"https://img.com/path/foo/bar.gif",
			},
		},
		{
			name: "ignore empty and invalid src",
			html: `<img src="">
				<img>No src</img>
				<img src="http://valid.com/img.jpg">
				<img src=" :badurl">`,
			base: base,
			expected: []string{
				"http://valid.com/img.jpg",
			},
		},
		{
			name:     "no images",
			html:     `<div>No images here!</div>`,
			base:     base,
			expected: []string{},
		},
		{
			name: "relative src with ..",
			html: `<img src="../up.jpg">`,
			base: base,
			expected: []string{
				"https://img.com/up.jpg",
			},
		},
		{
			name: "image with query and fragment",
			html: `<img src="/foo.png?bar=baz#frag">`,
			base: base,
			expected: []string{
				"https://img.com/foo.png?bar=baz#frag",
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getImagesFromHTML(tc.html, tc.base)
			if err != nil {
				t.Errorf("Test %v - %s FAIL: unexpected error: %v", i, tc.name, err)
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected %v, got %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
