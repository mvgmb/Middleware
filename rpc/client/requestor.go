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

var lookupOptions = &util.Options{
	Host:     "localhost",
	Port:     1337,
	Protocol: "tcp",
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

// Invoke works as the maestro
func (e *Requestor) Invoke(object string, req *proto.Message) (string, error) {
	serviceName := util.NewMessage([]byte(object), "Lookup", "OK", 200)

	result, err := e.Request(&serviceName, lookupOptions)
	if err != nil {
		return "", err
	}

	res, ok := result.(pb.Message)
	if !ok {
		return "", fmt.Errorf("Not a Message")
	}

	if res.Status.Code != 200 {
		return "", fmt.Errorf(res.Status.Message)
	}

	aor := util.StringToAOR(string(res.MessageData))

	options := util.Options{
		Host:     aor.Host,
		Port:     aor.Port,
		Protocol: "tcp",
	}

	result, err = e.Request(req, &options)
	if err != nil {
		return "", err
	}

	res, ok = result.(pb.Message)
	if !ok {
		return "", fmt.Errorf("Not a Message")
	}

	if res.Status.Code != 200 {
		return "", fmt.Errorf(res.Status.Message)
	}

	return string(res.MessageData), nil
}

// Request handles the entire proceso, from opening to closing one connection
func (e *Requestor) Request(req *proto.Message, options *util.Options) (interface{}, error) {
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
