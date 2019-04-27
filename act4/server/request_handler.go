package server

import (
	"fmt"
	"math"
	"net"

	"github.com/mvgmb/Middleware/act4/util"
)

// RequestHandler defines the ServerRequestHandler struct body
type RequestHandler struct {
	options  util.Options
	netConn  net.Conn
	listener net.Listener
}

// NewRequestHandler constructs a new ServerRequestHandler
func NewRequestHandler(options util.Options) (*RequestHandler, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", options.Port))
	if err != nil {
		return nil, err
	}

	e := &RequestHandler{
		options:  options,
		listener: listener,
	}

	return e, nil
}

// Accept accepts a new connection
func (e *RequestHandler) Accept() error {
	if e.netConn != nil {
		return fmt.Errorf("Already Accepted")
	}

	newConn, err := e.listener.Accept()
	if err != nil {
		return err
	}

	e.netConn = newConn

	return nil
}

// Close closes the connection with the client
func (e *RequestHandler) Close() error {
	if e.netConn == nil {
		return fmt.Errorf("Already Closed")
	}

	err := e.netConn.Close()
	if err != nil {
		return err
	}

	e.netConn = nil

	return nil
}

// Send sends a message to the defined server
func (e *RequestHandler) Send(message []byte) error {
	if e.netConn == nil {
		return fmt.Errorf("Not Accepted")
	}

	_, err := e.netConn.Write(message)
	return err
}

// Receive receives the response from the server
func (e *RequestHandler) Receive() ([]byte, error) {
	if e.netConn == nil {
		return nil, fmt.Errorf("Not Accepted")
	}

	buffer := make([]byte, math.MaxInt16)

	n, err := e.netConn.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer[:n], nil
}
