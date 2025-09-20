package main

import "testing"

func TestGetH1FromHTML(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name:     "simple h1",
			html:     "<html><body><h1>Test Title</h1></body></html>",
			expected: "Test Title",
		},
		{
			name:     "multiple h1s",
			html:     "<h1>First</h1><h1>Second</h1>",
			expected: "First",
		},
		{
			name:     "no h1 tag",
			html:     "<div>no heading</div>",
			expected: "",
		},
		{
			name:     "h1 with whitespace",
			html:     "<h1>   trim me   </h1>",
			expected: "trim me",
		},
		{
			name:     "h1 with nested tags",
			html:     "<h1>Welcome <span>World</span></h1>",
			expected: "Welcome World",
		},
		{
			name:     "empty string",
			html:     "",
			expected: "",
		},
		{
			name:     "malformed html",
			html:     "<h1>Broken",
			expected: "Broken",
		},
		{
			name:     "h1 not first tag",
			html:     "<div>before</div><h1>My Heading</h1>",
			expected: "My Heading",
		},
		{
			name:     "h1 with attributes",
			html:     `<h1 class="title">With Attr</h1>`,
			expected: "With Attr",
		},
		{
			name:     "h1 deeply nested",
			html:     `<div><header><h1>Deep</h1></header></div>`,
			expected: "Deep",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getH1FromHTML(tc.html)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected %q, got %q", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name: "main priority",
			html: `<html><body>
				<p>Outside paragraph.</p>
				<main>
					<p>Main paragraph.</p>
				</main>
			</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name: "no main, first p in body",
			html: `<html><body>
				<p>First para.</p>
				<p>Second para.</p>
			</body></html>`,
			expected: "First para.",
		},
		{
			name:     "no p in main, fallback to outside",
			html:     `<html><main></main><p>Fallback para</p></html>`,
			expected: "Fallback para",
		},
		{
			name:     "no p tag at all",
			html:     `<html><main><div>No paragraph</div></main></html>`,
			expected: "",
		},
		{
			name:     "first p nested in main",
			html:     `<main><div><p>Nested para</p></div></main>`,
			expected: "Nested para",
		},
		{
			name:     "malformed html",
			html:     `<main><p>Unclosed main`,
			expected: "Unclosed main",
		},
		{
			name:     "whitespace trimming",
			html:     `<main><p>   A  lot of   space   </p></main>`,
			expected: "A  lot of   space",
		},
		{
			name:     "p outside, no main at all",
			html:     `<p>Just one para</p>`,
			expected: "Just one para",
		},
		{
			name:     "main after p",
			html:     `<p>First para</p><main><p>In main</p></main>`,
			expected: "In main",
		},
		{
			name:     "main empty, p outside",
			html:     `<main></main><div><p>Outside main</p></div>`,
			expected: "Outside main",
		},
		{
			name:     "empty string",
			html:     "",
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.html)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected %q, got %q", i, tc.name, tc.expected, actual)
			}
		})
	}
}
