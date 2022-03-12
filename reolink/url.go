package reolink

import (
	"strings"
)

func getURL(address, path string) string {
	var url string
	if strings.HasPrefix(address, "http://") || strings.HasPrefix(address, "https://") {
		url = address
	} else {
		url = "http://" + address
	}

	if strings.HasSuffix(url, "/") {
		url = url[0 : len(url)-1]
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return url + path
}
