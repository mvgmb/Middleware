package util

import (
	"github.com/golang/protobuf/proto"
)

type Marshaller struct {
}

// NewMarshaller constructs a new Mashaller
func NewMarshaller() (*Marshaller, error) {
	return &Marshaller{}, nil
}

// Marshal serializes the message into bytes
func Marshal(message *proto.Message) ([]byte, error) {
	return proto.Marshal(*message)
}

// Unmarshal retrieves the serialized message
func Unmarshal(bytes *[]byte, pb proto.Message) error {
	err := proto.Unmarshal(*bytes, pb)
	if err != nil {
		return err
	}
	return nil
}

/*
proto.Message is an interface, Marshal/Unmarshal will receive &pb.MovieMessage{ ... }
https://developers.google.com/protocol-buffers/docs/gotutorial
*/
