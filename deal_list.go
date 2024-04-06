package gocheapshark

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/offerni/go-cheap-shark/errutils"
	"github.com/offerni/go-cheap-shark/utils"
)

func (c Client) DealList(opts DealListOpts) (*DealListResponse, error) {
	params := utils.BuildQueryParams(opts)

	jsonResp, err := c.call(callOpts{
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s/%s?%s", c.BaseURL, DealsAPIPath, params),
	})
	if err != nil {
		return nil, errutils.Wrap("c.call", err)
	}

	var deals []*DealFetchResponse
	if err := json.Unmarshal(jsonResp, &deals); err != nil {
		return nil, errutils.Wrap("json.Unmarshal(jsonResp, &deals)", err)
	}

	return &DealListResponse{
		Data: deals,
	}, nil
}

type DealListOpts struct {
	AAA         *bool   `json:"AAA"`
	Desc        *string `json:"desc"`
	Exact       *bool   `json:"exact"`
	LowerPrice  *uint   `json:"lowerPrice"`
	Metacritic  *uint   `json:"metacritic"`
	OnSale      *bool   `json:"onSale"`
	Output      *string `json:"output"`
	PageNumber  *uint   `json:"pageNumber"`
	PageSize    *uint   `json:"pageSize"`
	SortBy      *string `json:"sortBy"`
	SteamAppID  *string `json:"steamAppID"`
	SteamRating *uint   `json:"steamRating"`
	Steamworks  *bool   `json:"steamworks"`
	StoreID     *string `json:"storeID"`
	Title       *string `json:"title"`
	UpperPrice  *uint   `json:"upperPrice"`
}

type DealListResponse struct {
	Data []*DealFetchResponse `json:"data"`
}

type DealFetchResponse struct {
	DealID             string `json:"dealID"`
	DealRating         string `json:"dealRating"`
	GameID             string `json:"gameID"`
	InternalName       string `json:"internalName"`
	IsOnSale           string `json:"isOnSale"`
	LastChange         int64  `json:"lastChange"`
	MetacriticLink     string `json:"metacriticLink"`
	MetacriticScore    string `json:"metacriticScore"`
	NormalPrice        string `json:"normalPrice"`
	ReleaseDate        int64  `json:"releaseDate"`
	SalePrice          string `json:"salePrice"`
	Savings            string `json:"savings"`
	SteamAppID         string `json:"steamAppID"`
	SteamRatingCount   string `json:"steamRatingCount"`
	SteamRatingPercent string `json:"steamRatingPercent"`
	SteamRatingText    string `json:"steamRatingText"`
	StoreID            string `json:"storeID"`
	Thumb              string `json:"thumb"`
	Title              string `json:"title"`
}
