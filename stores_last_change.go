package gocheapshark

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/offerni/go-cheap-shark/errutils"
)

func (c Client) StoreLastChangeList() (*StoreLastChangeListResponse, error) {
	jsonResp, err := c.call(callOpts{
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s/%s?lastChange", c.BaseURL, StoresAPIPath),
	})
	if err != nil {
		return nil, errutils.Wrap("c.call", err)
	}

	var stores *StoreLastChangeFetchResponse
	if err := json.Unmarshal(jsonResp, &stores); err != nil {
		return nil, errutils.Wrap("json.Unmarshal(jsonResp, &stores)", err)
	}

	return &StoreLastChangeListResponse{
		Data: stores,
	}, nil
}

type StoreLastChangeFetchResponse map[string]string

type StoreLastChangeListResponse struct {
	Data *StoreLastChangeFetchResponse
}
