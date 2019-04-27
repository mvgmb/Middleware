package server

import (
	pb "github.com/mvgmb/Middleware/act4/proto"
	"github.com/mvgmb/Middleware/act4/util"
	"log"
)

// Invoker is the server side "maestro"
// Responsible for managing the invocations to the remote objects
type Invoker struct {
	requestHandler *RequestHandler
	marshaller     *util.Marshaller
	// serverProxy *ServerProxy
}

// NewInvoker constructs a new Invoker
func NewInvoker(options util.Options) (*Invoker, error) {
	rh, err := NewRequestHandler(options)
	if err != nil {
		return nil, err
	}

	marsh, err := util.NewMarshaller()
	if err != nil {
		return nil, err
	}

	// TODO create ServerProxy

	e := &Invoker{
		requestHandler: rh,
		marshaller:     marsh,
		//serverProxy: sp,
	}
	return e, nil
}

// Invoke is the core of the invoker
// Here is where he manage the clients requests
func (e *Invoker) Invoke() {
	// Create remote object

	for {
		// Invoke RequestHandler to wait for client message
		err := e.requestHandler.Accept()
		if err != nil {
			// TODO manage error
			log.Fatal(err)
		}

		// Receive request
		bytes, err := e.requestHandler.Receive()
		if err != nil {
			// TODO manage error
			log.Fatal(err)
		}

		// Unmarshal request
		req := pb.MovieMessage{}
		e.marshaller.Unmarshal(&bytes, &req)

		// test
		log.Println("Request", req.String())

		// Demultiplex where the message should go
		// TODO

		// Call required method
		// TODO

		// test
		res := util.NewMovieMessage([]byte("movie"), "BestPrice", "OK", 200)
		if err != nil {
			// TODO manage error
			log.Fatal(err)
		}

		// test
		log.Println("Responded", res.String())

		// Marshal the method output response
		bytes, err = e.marshaller.Marshal(&res)
		if err != nil {
			// TODO manage error
			log.Fatal(err)
		}

		// Invoke RequestHandler to send an answer to the client
		e.requestHandler.Send(bytes)

		// Close connection
		e.requestHandler.Close()
	}
}
