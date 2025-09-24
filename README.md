# Crawler

A fast and concurrent web crawler written in Go. This project crawls a website starting from a given URL, extracts key page information (such as the main heading, first paragraph, outgoing links, and images), and exports a report as a CSV file.

## Features

- **Configurable Concurrency:** Control how many pages are crawled at once.
- **Page Limits:** Set the maximum number of pages to crawl.
- **Data Extraction:** For each page, extract:
  - The normalized page URL
  - Main `<h1>` heading
  - First paragraph (`<p>`)
  - Outgoing link URLs
  - Image URLs
- **CSV Report:** Outputs all collected data to a CSV file.

## Usage

```sh
go build -o crawler
./crawler <URL> <maxConcurrency> <maxPages>
```

- `URL`: The starting point for the crawl (e.g., `https://example.com`)
- `maxConcurrency`: The maximum number of concurrent requests (e.g., `5`)
- `maxPages`: The maximum number of pages to crawl (e.g., `100`)

**Example:**

```sh
./crawler "https://www.wagslane.dev/" 5 100
```

This will crawl up to 100 pages from `https://www.wagslane.dev/` using 5 concurrent workers.

## Output

After the crawl is complete, a file named `report.csv` will be created in the current directory, containing the following columns:

- `page_url`: The URL of the crawled page
- `h1`: The main heading of the page
- `first_paragraph`: The first paragraph of the page
- `outgoing_link_urls`: Semicolon-separated list of outgoing link URLs from the page
- `image_urls`: Semicolon-separated list of image URLs from the page

## Example CSV Output

```csv
page_url,h1,first_paragraph,outgoing_link_urls,image_urls
https://example.com,Welcome,This is the first paragraph.,https://example.com/about;https://example.com/contact,https://example.com/img/logo.png
...
```

## Code Structure

- `main.go`: Entry point, argument parsing, crawler orchestration.
- `csv_report.go`: CSV report generation.
- Other files: Crawling logic and page data extraction.

## Building

Requires Go 1.18+.

```sh
go build -o crawler
```

## License

MIT License

---

*Made with Go 🕷️*