package reolink

import (
	"bytes"
	"encoding/json"
	"time"
)

func (c *Client) getToken() (string, error) {
	if c.token != "" && time.Now().Before(c.expiresAt) {
		return c.token, nil
	}

	reqBody, _ := json.Marshal([]Request{
		{
			Command: "Login",
			Action:  0,
			Param: Param{
				User: &User{
					Username: c.username,
					Password: c.password,
				},
			},
		},
	})

	resp, err := postRequest(
		getURL(c.address, "/cgi-bin/api.cgi?cmd=Login"),
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return "", err
	}

	token := resp.Value.Token.Name
	leaseTime := resp.Value.Token.LeaseTime
	c.token = token
	c.expiresAt = time.Now().Add(time.Second*leaseTime - 1*time.Minute)

	return token, nil
}
