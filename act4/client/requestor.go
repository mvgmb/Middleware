package client

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/mvgmb/Middleware/act4/proto"
	"github.com/mvgmb/Middleware/act4/util"
)

// Requestor deals with the access to the remote object
type Requestor struct {
	requestHandler *RequestHandler
	marshaller     *util.Marshaller
}

// NewRequestor constructs a new Requestor
func NewRequestor(options *util.Options) (*Requestor, error) {
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

// Invoke works as the maestro
func (e *Requestor) Invoke(options *util.Options, req *proto.Message) (interface{}, error) {
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

	return res, nil
}
