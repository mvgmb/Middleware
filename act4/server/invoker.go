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
func NewInvoker(options *util.Options) (*Invoker, error) {
	rh, err := NewRequestHandler(*options)
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

// Register registers a new service on the lookup table
func (e *Invoker) Register(serviceName string) error {
	options := util.Options{
		Host:     "localhost",
		Port:     1337,
		Protocol: "tcp",
	}
	err := e.requestHandler.Open(&options)
	if err != nil {
		return err
	}

	aor := util.AOR{
		Host:     e.requestHandler.options.Host,
		Port:     e.requestHandler.options.Port,
		ObjectID: serviceName,
	}

	log.Println(aor.String())

	req := util.NewMessage([]byte(aor.String()), "Bind", "OK", 200)
	bytes, err := e.marshaller.Marshal(&req)
	if err != nil {
		return err
	}

	err = e.requestHandler.Send(&bytes)
	if err != nil {
		return err
	}

	bytes, err = e.requestHandler.Receive()
	if err != nil {
		return err
	}

	res := pb.Message{}
	err = e.marshaller.Unmarshal(&bytes, &res)
	if err != nil {
		return err
	}

	if res.Status.Code == 200 {
		log.Println("noice")
	}

	err = e.requestHandler.Close()
	if err != nil {
		return err
	}

	return nil
}

// Invoke is the core of the invoker
// Here is where he manage the clients requests
func (e *Invoker) Invoke() {
	// Create remote object
	err := e.Register("BestPrice")
	if err != nil {
		log.Fatal(err)
	}

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
		req := pb.Message{}
		e.marshaller.Unmarshal(&bytes, &req)

		// test
		log.Println("Request", req.String())

		// Demultiplex where the message should go
		// TODO

		// Call required method
		// TODO

		// test
		res := util.NewMessage([]byte("movie"), "BestPrice", "OK", 200)
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
		e.requestHandler.Send(&bytes)

		// Close connection
		e.requestHandler.Close()
	}
}
