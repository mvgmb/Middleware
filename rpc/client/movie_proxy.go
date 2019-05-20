package client 

import (	
	pb "github.com/mvgmb/Middleware/rpc/proto"
	"github.com/mvgmb/Middleware/rpc/util"
	
	"fmt"
	"strconv"
)


// Proxy implements the MovieProxy
type Proxy struct {
	requestor *Requestor
}

var (
	options *util.Options
)

// NewProxy constructs a new MovieProxy
func NewProxy() (*Proxy, error) {
	requestor, err := NewRequestor()
	if err != nil {
		return nil, err
	}

	requestor.Lookup("Movie")

	e := &Proxy{
		requestor: requestor,
	}

	return e, nil
}

// MoviePrice return the requested movie price
func (e *Proxy) MoviePrice(movieName string) (int, error) {
	req := util.NewMessage([]byte(movieName), "Movie.Price", "OK", 200)

	res, err := e.requestor.Invoke(&req, options)
	if err != nil {
		return -1, err
	}

	message, ok := res.(*pb.Message)
	if !ok {
		return -1, fmt.Errorf("Not a Message")
	}

	num, err := strconv.ParseInt(string(message.MessageData), 10, 16)
	if err != nil {
		return -1, fmt.Errorf("Not a number")
	}

	return int(num), nil
}