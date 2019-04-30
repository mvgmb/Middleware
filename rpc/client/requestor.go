package client

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	pb "github.com/mvgmb/Middleware/rpc/proto"
	"github.com/mvgmb/Middleware/rpc/util"
)

// Requestor deals with the access to the remote object
type Requestor struct {
	requestHandler *RequestHandler
	marshaller     *util.Marshaller
}

// NewRequestor constructs a new Requestor
func NewRequestor() (*Requestor, error) {
	rh := NewRequestHandler()

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

// Invoke handles the entire proceso, from opening to closing one connection
func (e *Requestor) Invoke(req *proto.Message, options *util.Options) (proto.Message, error) {
	err := e.requestHandler.Open(*options)
	if err != nil {
		return nil, err
	}

	data, err := e.marshaller.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = e.requestHandler.Send(&data)
	if err != nil {
		return nil, err
	}

	data, err = e.requestHandler.Receive()
	if err != nil {
		return nil, err
	}

	res := pb.Message{}
	err = e.marshaller.Unmarshal(&data, &res)
	if err != nil {
		return nil, err
	}

	err = e.requestHandler.Close()
	if err != nil {
		return nil, err
	}

	if res.Status.Code != 200 {
		return nil, fmt.Errorf(res.Status.Message)
	}

	return &res, nil
}
