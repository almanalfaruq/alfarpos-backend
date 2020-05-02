package model

import "errors"

var (
	ErrNotFound error = errors.New("Not found in DB")
)
