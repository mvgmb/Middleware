package client 

import (	
	pb "github.com/mvgmb/Middleware/rpc/proto"
	"github.com/mvgmb/Middleware/rpc/util"
	
	"fmt"
	"strconv"
)


// Proxy implements the ClientProxy
type Proxy struct {
	requestor *Requestor
}

var (
	lookupOptions = &util.Options{
		Host:     "localhost",
		Port:     1337,
		Protocol: "tcp",
	}
	options *util.Options
)

// NewProxy constructs a new ClientProxy
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

// Lookup works as the maestro
func (e *Requestor) Lookup(object string) error {
	serviceName := util.NewMessage([]byte(object), "Lookup", "OK", 200)

	result, err := e.Invoke(&serviceName, lookupOptions)
	if err != nil {
		return err
	}

	res, ok := result.(*pb.Message)
	if !ok {
		return fmt.Errorf("Not a Message")
	}

	if res.Status.Code != 200 {
		return fmt.Errorf(res.Status.Message)
	}

	aor := util.StringToAOR(string(res.MessageData))

	options = &util.Options{
		Host:     aor.Host,
		Port:     aor.Port,
		Protocol: "tcp",
	}

	return nil
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

	num, err := strconv.ParseUint(string(message.MessageData), 10, 16)
	if err != nil {
		return -1, fmt.Errorf("Not a number")
	}

	return int(num), nil
}