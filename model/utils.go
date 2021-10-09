package model

import "errors"

var (
	ErrNotFound   error = errors.New("Not found in DB")
	ErrEmptyParam error = errors.New("Params should not be empty")
)
