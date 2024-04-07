package gocheapshark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/offerni/go-cheap-shark/errutils"
)

func (c Client) AlertManage(opts AlertManageOpts) (*string, error) {
	if err := opts.validate(); err != nil {
		return nil, err
	}

	if err := opts.validate(); err != nil {
		return nil, err
	}

	jsonResp, err := c.call(callOpts{
		Method: http.MethodGet,
		URL: fmt.Sprintf("%s/%s?action=%s&email=%s",
			c.BaseURL,
			AlertsAPIPath,
			opts.Action,
			opts.Email,
		)}) // again... too lazy (ﾉ˚Д˚)ﾉ
	if err != nil {
		return nil, errutils.Wrap("c.call", err)
	}

	if !json.Valid(jsonResp) {
		return nil, fmt.Errorf("%s", jsonResp)
	}

	var result string

	if bytes.Equal(jsonResp, []byte("[]")) {
		return &result, nil
	}

	err = json.Unmarshal(jsonResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (opts AlertManageOpts) validate() error {
	if opts.Action == "" {
		return ErrNoAction
	}

	if opts.Email == "" {
		return ErrNoEmail
	}

	return nil
}

type AlertManageOpts struct {
	Action string `json:"action"`
	Email  string `json:"email"`
}
