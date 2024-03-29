package server

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/mvgmb/Middleware/rpc/proto"
	"github.com/mvgmb/Middleware/rpc/util"

	"fmt"
	"log"
	"reflect"
	"strings"
)

// Invoker is the server side "maestro"
// Responsible for managing the invocations to the remote objects
type Invoker struct {
	requestHandler *RequestHandler
	marshaller     *util.Marshaller
	MovieProxy     *Proxy
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

	proxy := NewMovieProxy()

	e := &Invoker{
		requestHandler: rh,
		marshaller:     marsh,
		MovieProxy:     proxy,
	}
	return e, nil
}

// Invoke is the core of the invoker
// Here is where he manage the clients requests
func (e *Invoker) Invoke() {
	e.MovieProxy.NewMovieObject(e)

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

		var res proto.Message

		req := pb.Message{}
		err = e.marshaller.Unmarshal(&bytes, &req)
		if err != nil {
			log.Println(err)
			res = util.ErrBadRequest
		} else if req.Status.Code != 200 {
			res = util.ErrUnknown
		} else {
			call := strings.Split(req.TypeName, ".")
			args := strings.Split(string(req.MessageData), ",")

			switch call[0] {
			case "Movie":
				result := Call(e.MovieProxy.Movie, call[1], args[0])
				price := result[0].Int()
				res = util.NewMessage([]byte(fmt.Sprint(price)), "Price", "OK", 200)
			default:
				res = util.ErrBadRequest
			}
		}

		bytes, err = e.marshaller.Marshal(&res)
		if err != nil {
			log.Println(err)
		}

		e.requestHandler.Send(&bytes)
		if err != nil {
			log.Println(err)
		}

		e.requestHandler.Close()
		if err != nil {
			log.Println(err)
		}
	}
}

// Call calls a given method
func Call(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	return reflect.ValueOf(any).MethodByName(name).Call(inputs)
}
