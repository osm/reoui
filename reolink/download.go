package reolink

import (
	"fmt"
	"io"
)

func (c *Client) Download(file string, out io.Writer) error {
	token, err := c.getToken()
	if err != nil {
		return err
	}
	url := getURL(
		c.address,
		fmt.Sprintf("/cgi-bin/api.cgi?cmd=Download&source=%s&output=%s&token=%s",
			file, file, token,
		),
	)

	getRequest(url, out)
	return nil
}
