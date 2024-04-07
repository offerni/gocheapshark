package gocheapshark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/offerni/go-cheap-shark/errutils"
)

func (c Client) GameLookup(opts GameLookupOpts) (*GameLookupResponse, error) {
	if err := opts.validate(); err != nil {
		return nil, err
	}

	jsonResp, err := c.call(callOpts{
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s/%s?id=%d", c.BaseURL, GamesAPIPath, opts.ID),
	})
	if err != nil {
		return nil, errutils.Wrap("c.call", err)
	}

	if !json.Valid(jsonResp) {
		return nil, fmt.Errorf("%s", jsonResp)
	}

	if bytes.Equal(jsonResp, []byte("[]")) {
		return &GameLookupResponse{}, nil
	}

	var game GameLookupResponse
	if err := json.Unmarshal(jsonResp, &game); err != nil {
		return nil, errutils.Wrap("json.Unmarshal(jsonResp, &game)", err)
	}

	// Adding the requested ID back to the payload for consistency since it's not
	// returned by default
	game.Info.GameID = &opts.ID

	return &game, nil
}

func (opts GameLookupOpts) validate() error {
	if opts.ID == 0 {
		return ErrNoID
	}

	return nil
}

type GameLookupOpts struct {
	ID uint `json:"id"`
}

type GameLookupResponse struct {
	CheapestPriceEver *CheapestPrice    `json:"cheapestPriceEver"`
	Deals             []*GameLookupDeal `json:"deals"`
	Info              *GameLookupInfo   `json:"info"`
}

type GameLookupInfo struct {
	GameID *uint   `json:"gameID"`
	Thumb  *string `json:"thumb"`
	Title  *string `json:"title"`
}

type GameLookupDeal struct {
	StoreID     *string `json:"storeID"`
	DealID      *string `json:"dealID"`
	Price       *string `json:"price"`
	RetailPrice *string `json:"retailPrice"`
	Savings     *string `json:"savings"`
}
