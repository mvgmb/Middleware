package client 

import (	
	"github.com/mvgmb/Middleware/rpc/util"
	"fmt"
	"strconv"
)


// Proxy implements the ClientProxy
type Proxy struct {
	requestor *Requestor
}

// NewProxy constructs a new ClientProxy
func NewProxy() (*Proxy, error) {
	requestor, err := NewRequestor()
	if err != nil {
		return nil, err
	}

	e := &Proxy{
		requestor: requestor,
	}

	return e, nil
}

// MoviePrice return the requested movie price
func (e *Proxy) MoviePrice(movieName string) (int, error) {
	req := util.NewMessage([]byte(movieName), "Movie.Price", "OK", 200)

	res, err := e.requestor.Invoke("Movie", &req)
	if err != nil {
		return -1, err
	}
	
	num, err := strconv.ParseUint(res, 10, 16)
	if err != nil {
		return -1, fmt.Errorf("Not a number")
	}

	return int(num), nil
}