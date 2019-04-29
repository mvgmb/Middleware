package util

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/mvgmb/Middleware/act4/proto"
)

var (
	ErrUnknown            = NewMessage([]byte(""), "", "Unknown", 000)
	ErrBadRequest         = NewMessage([]byte(""), "", "Bad Request", 400)
	ErrUnauthorized       = NewMessage([]byte(""), "", "Unauthorized", 401)
	ErrForbidden          = NewMessage([]byte(""), "", "Forbidden", 403)
	ErrNotFound           = NewMessage([]byte(""), "", "Service not found", 404)
	ErrMethodNotAllowed   = NewMessage([]byte(""), "", "Method not allowed", 405)
	ErrPayloadTooLarge    = NewMessage([]byte(""), "", "Payload too large!", 413)
	ErrExpectationFailed  = NewMessage([]byte(""), "", "Expectation fail", 417)
	ErrServiceUnavailable = NewMessage([]byte(""), "", "Service unavailable", 503)
)

// Options defines the options values
type Options struct {
	Host     string
	Port     uint16
	Protocol string
}

// NewMessage creates a new generic Message instance
func NewMessage(bytes []byte, typeName, statusMessage string, statusCode uint64) proto.Message {
	status := &pb.Status{
		Code:    statusCode,
		Message: statusMessage,
	}
	message := &pb.Message{
		TypeName:    typeName,
		MessageData: bytes,
		Status:      status,
	}
	return message
}
