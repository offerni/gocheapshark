package gocheapshark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/offerni/go-cheap-shark/errutils"
)

func (c Client) AlertEdit(opts AlertEditOpts) (*bool, error) {
	if err := opts.validate(); err != nil {
		return nil, err
	}

	if err := opts.validate(); err != nil {
		return nil, err
	}

	jsonResp, err := c.call(callOpts{
		Method: http.MethodGet,
		URL: fmt.Sprintf("%s/%s?action=%s&email=%s&gameID=%d&price=%.2f",
			c.BaseURL,
			AlertsAPIPath,
			opts.Action,
			opts.Email,
			opts.GameID,
			opts.Price,
		)}) // yeah, I know... I was a little lazy on this one (ﾉ˚Д˚)ﾉ
	if err != nil {
		return nil, errutils.Wrap("c.call", err)
	}

	if !json.Valid(jsonResp) {
		return nil, fmt.Errorf("%s", jsonResp)
	}

	var result bool

	if bytes.Equal(jsonResp, []byte("[]")) {
		return &result, nil
	}

	err = json.Unmarshal(jsonResp, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (opts AlertEditOpts) validate() error {
	if opts.Action == "" {
		return ErrNoAction
	}

	if opts.Email == "" {
		return ErrNoEmail
	}

	if opts.GameID == 0 {
		return ErrNoID
	}

	return nil
}

type AlertEditOpts struct {
	Action string  `json:"action"`
	Email  string  `json:"email"`
	GameID uint    `json:"gameID"`
	Price  float64 `json:"price"`
}
