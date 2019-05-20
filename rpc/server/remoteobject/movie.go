package remoteobject

import "fmt"

// Movie struct
type Movie struct{}

var movies = map[string]int{
	"Titanic":       23,
	"Lagoa Azul":    14,
	"Lilo e Stitch": 17,
	"Matilda":       10,
}

// Price returns the movie price
func (e *Movie) Price(movieName string) (int, error) {
	if _, ok := movies[movieName]; !ok {
		return -1, fmt.Errorf("Movie not found")
	}

	return movies[movieName], nil
}
