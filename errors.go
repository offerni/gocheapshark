package gocheapshark

import "errors"

var (
	ErrNoAction  = errors.New("Action is required")
	ErrNoBaseURL = errors.New("Base URL is required")
	ErrNoEmail   = errors.New("Email is required")
	ErrNoTitle   = errors.New("Title is required")
	ErrNoID      = errors.New("ID is required")
	ErrNoIDs     = errors.New("At least one ID is required")
	ErrNoPrice   = errors.New("Price is required")
)
