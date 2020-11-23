package muxxxer

import (
	"net/http"
	"net/url"
	"strings"
)

type dispatcher struct {
	f func(http.ResponseWriter, *http.Request)
}

func (d *dispatcher) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	d.f(rw, r)
}

// Route is a combination of an http.Handler and the path it is registered to handle
type Route struct {
	Path    string
	Handler http.Handler
}

// NewRoute creates a new Route by converting the handlerFunc to a Handler
func NewRoute(p string, f func(http.ResponseWriter, *http.Request)) *Route {
	return &Route{
		p,
		&dispatcher{f},
	}
}

func (r *Route) handles(uri *url.URL) bool {

	if strings.HasSuffix(r.Path, "/") {
		return strings.HasPrefix(
			strings.TrimLeft(uri.Path, "/"),
			strings.Trim(r.Path, "/"))
	} else {
		return strings.Trim(uri.Path, "/") == strings.Trim(r.Path, "/")
	}
}
