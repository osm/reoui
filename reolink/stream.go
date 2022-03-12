package reolink

import (
	"fmt"
	"net/http"
)

func (c *Client) Stream(w http.ResponseWriter) (*http.Response, error) {
	token, err := c.getToken()
	if err != nil {
		return nil, err
	}

	quality := "main"
	if c.lowStreamQuality == true {
		quality = "sub"
	}
	url := getURL(
		c.address,
		fmt.Sprintf("flv?port=1935&app=bcs&stream=channel0_%s.bcs&token=%s", quality, token),
	)

	return http.Get(url)
}
