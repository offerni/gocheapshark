package gocheapshark

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

type callOpts struct {
	Body   *bytes.Buffer
	Method string
	URL    string
}

func (c *Client) call(opts callOpts) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Duration(time.Second * 10),
		Transport: &http.Transport{
			MaxIdleConns:        1000,
			MaxIdleConnsPerHost: 1000,
		},
	}
	req, err := http.NewRequest(opts.Method, opts.URL, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
