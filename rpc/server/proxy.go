package server

import (
	"github.com/mvgmb/Middleware/rpc/movie"
)

// Proxy declares the ServerProxy
type Proxy struct {
	Movie *movie.Movie
}

// NewProxy constructs a new ServerProxy
func NewProxy() *Proxy {
	return &Proxy{}
}

// NewMovieObject constructs a new Moview object instance
func (e *Proxy) NewMovieObject(invoker *Invoker) error {
	err := invoker.Register("Movie")
	if err != nil {
		return err
	}

	e.Movie = &movie.Movie{}
	return nil
}
