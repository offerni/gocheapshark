package gocheapshark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/offerni/gocheapshark/errutils"
)

func (c Client) StoreList() (*StoreListResponse, error) {
	jsonResp, err := c.call(callOpts{
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s/%s", c.BaseURL, StoresAPIPath),
	})
	if err != nil {
		return nil, errutils.Wrap("c.call", err)
	}

	if !json.Valid(jsonResp) {
		return nil, fmt.Errorf("%s", jsonResp)
	}

	if bytes.Equal(jsonResp, []byte("[]")) {
		return &StoreListResponse{}, nil
	}

	var stores []*StoreFetchResponse
	if err := json.Unmarshal(jsonResp, &stores); err != nil {
		return nil, errutils.Wrap("json.Unmarshal(jsonResp, &stores)", err)
	}

	return &StoreListResponse{
		Data: stores,
	}, nil
}

type StoreListResponse struct {
	Data []*StoreFetchResponse `json:"data"`
}

type StoreFetchResponse struct {
	Images    *StoreImages `json:"images"`
	IsActive  *uint        `json:"isActive"`
	StoreID   *string      `json:"storeID"`
	StoreName *string      `json:"storeName"`
}

type StoreImages struct {
	Banner *string `json:"banner"`
	Icon   *string `json:"icon"`
	Logo   *string `json:"logo"`
}
