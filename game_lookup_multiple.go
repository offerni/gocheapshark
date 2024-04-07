package gocheapshark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/offerni/go-cheap-shark/errutils"
	"github.com/offerni/go-cheap-shark/utils"
)

func (c Client) GameLookupMultiple(opts GameLookupMultipleOpts) (*GameLookupMultipleResponse, error) {
	if err := opts.validate(); err != nil {
		return nil, err
	}

	ids := strings.Join(opts.IDs, ",")

	jsonResp, err := c.call(callOpts{
		Method: http.MethodGet,
		URL:    fmt.Sprintf("%s/%s?ids=%s", c.BaseURL, GamesAPIPath, ids),
	})
	if err != nil {
		return nil, errutils.Wrap("c.call", err)
	}

	if !json.Valid(jsonResp) {
		return nil, fmt.Errorf("%s", jsonResp)
	}

	if bytes.Equal(jsonResp, []byte("[]")) {
		return &GameLookupMultipleResponse{}, nil
	}

	var games GameLookupMultipleResponse
	if err := json.Unmarshal(jsonResp, &games); err != nil {
		return nil, errutils.Wrap("json.Unmarshal(jsonResp, &game)", err)
	}

	// Adding the requested IDs back to the payload under info for consistency
	// since it's not returned by default
	remappeGames := make(GameLookupMultipleResponse, len(games))
	for i, game := range games {
		id, err := utils.ConvertNumericStringToUintPointer(i)
		if err != nil {
			return nil, errutils.Wrap("utils.ConvertNumericStringToUintPointer", err)
		}

		game.Info.GameID = id

		remappeGames[i] = game
	}

	return &remappeGames, nil
}

func (opts GameLookupMultipleOpts) validate() error {
	if len(opts.IDs) == 0 {
		return ErrNoIDs
	}

	return nil
}

type GameLookupMultipleOpts struct {
	IDs []string `json:"ids"`
}

type GameLookupMultipleResponse map[string]GameLookupResponse
