package client

import (
	"fmt"
	"math"
	"net"

	"github.com/mvgmb/Middleware/act4/util"
)

// RequestHandler implements the client side communications manager
type RequestHandler struct {
	options util.Options
	netConn net.Conn
}

// NewRequestHandler declares a new RequestHandler
func NewRequestHandler(options util.Options) (*RequestHandler, error) {
	if options.Protocol != "tcp" && options.Protocol != "udp" {
		return nil, util.ErrMethodNotAllowed
	}

	netConn, err := net.Dial(options.Protocol, fmt.Sprintf("%s:%d", options.Host, options.Port))
	if err != nil {
		return nil, err
	}

	e := &RequestHandler{
		options: options,
		netConn: netConn,
	}

	return e, nil
}

// Send sends a message to the defined server
func (e *RequestHandler) Send(message *[]byte) error {
	_, err := e.netConn.Write(*message)
	return err
}

// Receive receives the response from the server
func (e *RequestHandler) Receive() ([]byte, error) {
	buffer := make([]byte, math.MaxInt16)

	n, err := e.netConn.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer[:n], nil
}

// Close closes the connection with the server
func (e *RequestHandler) Close() error {
	return e.netConn.Close()
}
