package util

import (
	"errors"
	"github.com/golang/protobuf/proto"
	pb "github.com/mvgmb/Middleware/act4/proto"
)

var (
	ErrUnknown            = errors.New("000 - Unknown")
	ErrUnauthorized       = errors.New("401 - Unauthorized")
	ErrForbidden          = errors.New("403 - Forbidden")
	ErrNotFound           = errors.New("404 - Not Found")
	ErrMethodNotAllowed   = errors.New("405 - Method Not Allowed")
	ErrPayloadTooLarge    = errors.New("413 - Payload Too Large")
	ErrExpectationFailed  = errors.New("417 - Expectation Failed")
	ErrServiceUnavailable = errors.New("503 - Service Unavailable")
)

// Options defines the options values
type Options struct {
	Host     string
	Port     uint16
	Protocol string
}

// NewMovieMessage creates a new MovieMessage instance
func NewMovieMessage(bytes []byte, typeName, statusMessage string, statusCode uint64) proto.Message {
	status := &pb.Status{
		Code:    statusCode,
		Message: statusMessage,
	}
	message := &pb.MovieMessage{
		TypeName:    typeName,
		MessageData: bytes,
		Status:      status,
	}
	return message
}
