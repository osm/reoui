package reolink

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func postRequest(url string, payload *bytes.Buffer) (*Response, error) {
	resp, err := http.Post(
		url,
		"application/json",
		payload,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ret []Response
	if err := json.Unmarshal(respBody, &ret); err != nil {
		return nil, fmt.Errorf("unable to unmarshal json, %v, raw data: %v", err, string(respBody))
	}

	if len(ret) != 1 {
		return nil, fmt.Errorf("unknown response, %v", ret)
	}
	if ret[0].Error != nil {
		return nil, fmt.Errorf("error: %v", ret[0].Error.Detail)
	}

	return &ret[0], nil
}

func getRequest(url string, out io.Writer) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
