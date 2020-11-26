package muxxxer

import "net/http"

// ArgumentBag contains ways to get path and query args from the matching url.
type ArgumentBag struct{}

func newArgumentBag(r string, req *http.Request) *ArgumentBag {
	return &ArgumentBag{}
}

// NewArgumentRoute creates a route with access to an ArgumentBag
func NewArgumentRoute(r string, f func(http.ResponseWriter, *http.Request, *ArgumentBag)) *Route {
	return NewRoute(r, func(rw http.ResponseWriter, req *http.Request) {
		f(rw, req, newArgumentBag(r, req))
	})
}
