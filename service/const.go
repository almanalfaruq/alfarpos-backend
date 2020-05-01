package service

import "errors"

var (
	ErrEmptyParam error = errors.New("Params should not be empty")
)
