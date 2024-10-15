package storage

import "errors"

var (
	URLNotFound = errors.New("URL not found")
	URLExists   = errors.New("URL exists")
)
