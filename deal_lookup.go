package gocheapshark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/offerni/go-cheap-shark/errutils"
)

func (c Client) DealLookup(opts DealLookupOpts) (*DealLookupResponse, error) {
	if err := opts.validate(); err != nil {
		return nil, err
	}

	jsonResp, err := c.call(callOpts{
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s/%s?id=%s", c.BaseURL, DealsAPIPath, opts.ID),
	})
	if err != nil {
		return nil, errutils.Wrap("c.call", err)
	}

	if !json.Valid(jsonResp) {
		return nil, fmt.Errorf("%s", jsonResp)
	}

	if bytes.Equal(jsonResp, []byte("[]")) {
		return &DealLookupResponse{}, nil
	}

	var deal DealLookupResponse
	if err := json.Unmarshal(jsonResp, &deal); err != nil {
		return nil, errutils.Wrap("json.Unmarshal(jsonResp, &deal)", err)
	}

	// Adding the requested ID back to the payload for consistency since it's not
	// returned by default
	deal.GameInfo.DealID = &opts.ID

	return &deal, nil
}

func (opts DealLookupOpts) validate() error {
	if opts.ID == "" {
		return ErrNoID
	}

	return nil
}

type DealLookupOpts struct {
	ID string `json:"id"`
}

type DealLookupResponse struct {
	CheaperStores []*CheaperStore `json:"cheaperStores"`
	CheapestPrice *CheapestPrice  `json:"cheapestPrice"`
	GameInfo      *GameInfo       `json:"gameInfo"`
}

type GameInfo struct {
	DealID             *string `json:"dealID"`
	GameID             *string `json:"gameID"`
	MetacriticLink     *string `json:"metacriticLink"`
	MetacriticScore    *string `json:"metacriticScore"`
	Name               *string `json:"name"`
	Publisher          *string `json:"publisher"`
	ReleaseDate        *int64  `json:"releaseDate"`
	RetailPrice        *string `json:"retailPrice"`
	SalePrice          *string `json:"salePrice"`
	SteamAppID         *string `json:"steamAppID"`
	SteamRatingCount   *string `json:"steamRatingCount"`
	SteamRatingPercent *string `json:"steamRatingPercent"`
	SteamRatingText    *string `json:"steamRatingText"`
	Steamworks         *string `json:"steamworks"`
	StoreID            *string `json:"storeID"`
	Thumb              *string `json:"thumb"`
}

type CheaperStore struct {
	DealID      string `json:"dealID"`
	RetailPrice string `json:"retailPrice"`
	SalePrice   string `json:"salePrice"`
	StoreID     string `json:"storeID"`
}

type CheapestPrice struct {
	Date  int64  `json:"date"`
	Price string `json:"price"`
}
