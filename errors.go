package gocheapshark

import "errors"

var (
	ErrNoBaseURL   = errors.New("Base URL is required")
	ErrNoID        = errors.New("ID is required")
	ErrNoIDs       = errors.New("At least one ID is required")
	ErrNoGameTitle = errors.New("Game Title is required")
)
