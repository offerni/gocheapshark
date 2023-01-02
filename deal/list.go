package deal

import (
	"time"
)

type ListOpts struct {
	AAA         *bool   `json:"AAA"`
	Desc        *string `json:"desc"`
	Exact       *bool   `json:"exact"`
	LowerPrice  *uint   `json:"lowerPrice"`
	Metacritic  *uint   `json:"metacritic"`
	OnSale      *bool   `json:"onSale"`
	Output      *string `json:"output"`
	PageNumber  *uint   `json:"pageNumber"`
	PageSize    *uint   `json:"pageSize"`
	SortBy      *string `json:"sortBy"` // FUTURE: create a custom type with all the possible values
	SteamAppID  *string `json:"steamAppID"`
	SteamRating *uint   `json:"steamRating"`
	Steamworks  *bool   `json:"steamworks"`
	StoreID     *string `json:"storeID"`
	Title       *string `json:"title"`
	UpperPrice  *uint   `json:"upperPrice"`
}

type ListResponse struct {
	Data []*Response
}

type Response struct {
	DealID             string    `json:"dealID"`
	DealRating         string    `json:"dealRating"`
	GameID             string    `json:"gameID"`
	InternalName       string    `json:"internalName"`
	IsOnSale           string    `json:"isOnSale"`
	LastChange         time.Time `json:"lastChange"`
	MetacriticLink     string    `json:"metacriticLink"`
	MetacriticScore    string    `json:"metacriticScore"`
	NormalPrice        string    `json:"normalPrice"`
	ReleaseDate        time.Time `json:"releaseDate"`
	SalePrice          string    `json:"salePrice"`
	Savings            string    `json:"savings"`
	SteamAppID         string    `json:"steamAppID"`
	SteamRatingCount   string    `json:"steamRatingCount"`
	SteamRatingPercent string    `json:"steamRatingPercent"`
	SteamRatingText    string    `json:"steamRatingText"`
	StoreID            string    `json:"storeID"`
	Thumb              string    `json:"thumb"`
	Title              string    `json:"title"`
}
