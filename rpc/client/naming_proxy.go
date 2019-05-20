package client

import (
	"fmt"

	pb "github.com/mvgmb/Middleware/rpc/proto"
	"github.com/mvgmb/Middleware/rpc/util"
)

var lookupOptions = &util.Options{
	Host:     "localhost",
	Port:     1337,
	Protocol: "tcp",
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
