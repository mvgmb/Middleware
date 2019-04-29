package client

import (
	"errors"
	"fmt"
	"math"
	"net"

	"github.com/mvgmb/Middleware/rpc/util"
)

// RequestHandler implements the client side communications manager
type RequestHandler struct {
	netConn net.Conn
}

// NewRequestHandler declares a new RequestHandler
func NewRequestHandler() *RequestHandler {
	e := RequestHandler{
		netConn: nil,
	}
	return &e
}

// Open opens a new connection
func (e *RequestHandler) Open(options util.Options) error {
	if e.netConn != nil {
		return errors.New("Connection established, please close to open a new one")
	}

	netConn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", options.Host, options.Port))
	if err != nil {
		return err
	}
	e.netConn = netConn
	return nil
}

// Close closes the connection with the server
func (e *RequestHandler) Close() error {
	if e.netConn == nil {
		return errors.New("No connection established")
	}

	err := e.netConn.Close()
	if err != nil {
		return err
	}
	e.netConn = nil
	return nil
}

// Send sends a message to the defined server
func (e *RequestHandler) Send(message *[]byte) error {
	if e.netConn == nil {
		return errors.New("No connection established")
	}
	_, err := e.netConn.Write(*message)
	return err
}

// Receive receives the response from the server
func (e *RequestHandler) Receive() ([]byte, error) {
	if e.netConn == nil {
		return nil, errors.New("No connection established")
	}

	buffer := make([]byte, math.MaxInt16)

	n, err := e.netConn.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer[:n], nil
}
