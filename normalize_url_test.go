package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove http scheme",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove trailing slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "both http and trailing slash",
			inputURL: "http://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "no path, no trailing slash",
			inputURL: "https://blog.boot.dev",
			expected: "blog.boot.dev",
		},
		{
			name:     "no path, trailing slash",
			inputURL: "https://blog.boot.dev/",
			expected: "blog.boot.dev",
		},
		{
			name:     "www prefix",
			inputURL: "https://www.blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "www prefix and trailing slash",
			inputURL: "http://www.blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "www prefix, no path",
			inputURL: "https://www.blog.boot.dev",
			expected: "blog.boot.dev",
		},
		{
			name:     "www prefix, no path, trailing slash",
			inputURL: "https://www.blog.boot.dev/",
			expected: "blog.boot.dev",
		},
		{
			name:     "path with multiple segments",
			inputURL: "https://blog.boot.dev/a/b/c",
			expected: "blog.boot.dev/a/b/c",
		},
		{
			name:     "path with trailing slash, multiple segments",
			inputURL: "https://blog.boot.dev/a/b/c/",
			expected: "blog.boot.dev/a/b/c",
		},
		{
			name:     "www prefix, multiple path segments with trailing slash",
			inputURL: "http://www.blog.boot.dev/a/b/c/",
			expected: "blog.boot.dev/a/b/c",
		},
		{
			name:     "query string is ignored",
			inputURL: "https://blog.boot.dev/path?query=123",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "fragment is ignored",
			inputURL: "https://blog.boot.dev/path#section2",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "query and fragment ignored",
			inputURL: "https://blog.boot.dev/path/?q=1#frag",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "path is just slash",
			inputURL: "https://blog.boot.dev/",
			expected: "blog.boot.dev",
		},
		{
			name:     "root path with www",
			inputURL: "http://www.blog.boot.dev/",
			expected: "blog.boot.dev",
		},
		{
			name:     "localhost with port",
			inputURL: "http://localhost:8080/foo/bar/",
			expected: "localhost:8080/foo/bar",
		},
		{
			name:     "domain with port and trailing slash",
			inputURL: "https://example.com:3000/test/",
			expected: "example.com:3000/test",
		},
		{
			name:     "domain with port, no path",
			inputURL: "https://example.com:3000",
			expected: "example.com:3000",
		},
		{
			name:     "double slash in path",
			inputURL: "https://blog.boot.dev//double//slash/",
			expected: "blog.boot.dev//double//slash",
		},
		{
			name:     "empty path",
			inputURL: "https://blog.boot.dev",
			expected: "blog.boot.dev",
		},
		{
			name:     "path with leading slash",
			inputURL: "https://blog.boot.dev//path",
			expected: "blog.boot.dev//path",
		},
		{
			name:     "path with only slash",
			inputURL: "https://blog.boot.dev/",
			expected: "blog.boot.dev",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
