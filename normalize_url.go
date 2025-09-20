package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawurl string) (string, error) {
	u, err := url.Parse(rawurl)

	if err != nil {
		return "", err
	}

	host := u.Host
	host = strings.TrimPrefix(host, "www.")
	path := u.Path

	if path != "/" && strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}

	path = strings.TrimPrefix(path, "/")
	normalized := host

	if path != "" {
		normalized += "/" + path
	}

	return normalized, nil
}
