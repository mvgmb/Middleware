package server

import "github.com/mvgmb/Middleware/rpc/server/remoteobject"

// Proxy declares the ServerProxy
type Proxy struct {
	Movie *remoteobject.Movie
}

// NewProxy constructs a new ServerProxy
func NewMovieProxy() *Proxy {
	return &Proxy{}
}

// NewMovieObject constructs a new Moview object instance
func (e *Proxy) NewMovieObject(invoker *Invoker) error {
	err := invoker.Register("Movie")
	if err != nil {
		return err
	}

	e.Movie = &remoteobject.Movie{}
	return nil
}
