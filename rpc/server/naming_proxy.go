package server

import (
	"github.com/mvgmb/Middleware/rpc/client"
	"github.com/mvgmb/Middleware/rpc/util"
)

// Register registers a new object on the lookup table
func (e *Invoker) Register(remoteObjectName string) error {
	options := util.Options{
		Host:     "localhost",
		Port:     1337,
		Protocol: "tcp",
	}

	clientRequestor, err := client.NewRequestor()
	if err != nil {
		return err
	}

	aor := util.AOR{
		Host:     e.requestHandler.options.Host,
		Port:     e.requestHandler.options.Port,
		ObjectID: remoteObjectName,
	}

	req := util.NewMessage([]byte(aor.String()), "Bind", "OK", 200)

	clientRequestor.Invoke(&req, &options)

	return nil
}
