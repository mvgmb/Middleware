package util

import "errors"

var (
	ErrMethodNotAllowed = errors.New("405 - Method Not Allowed")
)

// Options defines the options values
type Options struct {
	Host     string
	Port     uint16
	Protocol string
}
