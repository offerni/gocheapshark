package gocheapshark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/offerni/gocheapshark/errutils"
	"github.com/offerni/gocheapshark/utils"
)

func (c Client) GameList(opts GameListOpts) (*GameListResponse, error) {
	if err := opts.validate(); err != nil {
		return nil, err
	}

	params := utils.BuildQueryParams(opts)

	jsonResp, err := c.call(callOpts{
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s/%s?%s", c.BaseURL, GamesAPIPath, params),
	})
	if err != nil {
		return nil, errutils.Wrap("c.call", err)
	}

	if !json.Valid(jsonResp) {
		return nil, fmt.Errorf("%s", jsonResp)
	}

	if bytes.Equal(jsonResp, []byte("[]")) {
		return &GameListResponse{}, nil
	}

	var games []*GameFetchResponse
	if err := json.Unmarshal(jsonResp, &games); err != nil {
		return nil, errutils.Wrap("json.Unmarshal(jsonResp, &games)", err)
	}

	return &GameListResponse{
		Data: games,
	}, nil
}

func (opts GameListOpts) validate() error {
	if opts.Title == nil {
		return ErrNoTitle
	}

	return nil
}

type GameListOpts struct {
	Title      *string `json:"title"`
	SteamAppID *uint   `json:"steamAppID"`
	Limit      *uint   `json:"limit"`
	Exact      *bool   `json:"exact"`
}

type GameListResponse struct {
	Data []*GameFetchResponse `json:"data"`
}

type GameFetchResponse struct {
	GameID         *string `json:"gameID"`
	SteamAppID     *string `json:"steamAppID"`
	Cheapest       *string `json:"cheapest"`
	CheapestDealID *string `json:"cheapestDealID"`
	External       *string `json:"external"`
	InternalName   *string `json:"internalName"`
	Thumb          *string `json:"thumb"`
}
