package client

import (
	"github.com/golang/protobuf/proto"
	"github.com/mvgmb/Middleware/act4/util"
)

// Requestor deals with the access to the remote object
type Requestor struct {
	requestHandler *RequestHandler
	marshaller     *util.Marshaller
}

// NewRequestor constructs a new Requestor
func NewRequestor(options util.Options) (*Requestor, error) {
	rh, err := NewRequestHandler(options)
	if err != nil {
		return nil, err
	}
	marsh, err := util.NewMarshaller()
	if err != nil {
		return nil, err
	}

	e := &Requestor{
		requestHandler: rh,
		marshaller:     marsh,
	}

	return e, nil
}

// Close closes the request handler connection
func (e *Requestor) Close() error {
	return e.requestHandler.Close()
}

// Invoke works as the maestro
func (e *Requestor) Invoke(req *proto.Message, res proto.Message) error {
	// serialize
	data, err := e.marshaller.Marshal(req)
	if err != nil {
		return err
	}

	// send
	err = e.requestHandler.Send(data)
	if err != nil {
		return err
	}

	// receive
	data, err = e.requestHandler.Receive()
	if err != nil {
		return err
	}

	// deserialize and sends to client proxy
	err = e.marshaller.Unmarshal(&data, res)
	if err != nil {
		return err
	}

	return nil
}
