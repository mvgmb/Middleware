package server

import (
	"github.com/mvgmb/Middleware/rpc/movie"
)

// Proxy declares the ServerProxy
type Proxy struct {
	Invoker *Invoker
	Movie   *movie.Movie
}

// NewProxy constructs a new ServerProxy
func NewProxy() *Proxy {
	return &Proxy{}
}

// NewMovieObject constructs a new Moview object instance
func (e *Proxy) NewMovieObject() {
	e.Movie = &movie.Movie{}
}
