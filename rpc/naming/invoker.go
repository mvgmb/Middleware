package naming

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/mvgmb/Middleware/rpc/proto"
	"github.com/mvgmb/Middleware/rpc/util"

	"log"
)

// Invoker is the server side "maestro"
// Responsible for managing the invocations to the remote objects
type Invoker struct {
	requestHandler *RequestHandler
	marshaller     *util.Marshaller
}

// NewInvoker constructs a new Invoker
func NewInvoker(options *util.Options) (*Invoker, error) {
	rh, err := NewRequestHandler(*options)
	if err != nil {
		return nil, err
	}

	marsh, err := util.NewMarshaller()
	if err != nil {
		return nil, err
	}

	e := &Invoker{
		requestHandler: rh,
		marshaller:     marsh,
	}
	return e, nil
}

// Invoke is the core of the invoker
// Here is where he manage the requests
func (e *Invoker) Invoke() {
	log.Printf("Listening at %s:%d\n", e.requestHandler.options.Host, e.requestHandler.options.Port)

	for {
		err := e.requestHandler.Accept()
		if err != nil {
			log.Println(err)
			e.requestHandler.Close()
			continue
		}

		bytes, err := e.requestHandler.Receive()
		if err != nil {
			log.Println(err)
			e.requestHandler.Close()
			continue
		}

		var message proto.Message

		req := pb.Message{}

		err = e.marshaller.Unmarshal(&bytes, &req)
		if err != nil {
			message = util.ErrBadRequest
		} else {
			switch req.TypeName {
			case "Lookup":
				result, err := lookup(string(req.MessageData))
				if err != nil {
					message = util.ErrNotFound
					break
				}
				message = util.NewMessage([]byte(result.String()), "AOR", "OK", 200)
			case "Bind":
				bind(util.StringToAOR(string(req.MessageData)))
				message = util.NewMessage([]byte(""), "", "OK", 200)
			default:
				message = util.ErrBadRequest
			}
		}

		bytes, err = e.marshaller.Marshal(&message)
		if err != nil {
			log.Println(err)
			e.requestHandler.Close()
			continue
		}

		err = e.requestHandler.Send(&bytes)
		if err != nil {
			log.Println(err)
			e.requestHandler.Close()
			continue
		}

		err = e.requestHandler.Close()
		if err != nil {
			log.Println(err)
		}
	}
}
