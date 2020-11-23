package muxxxer

import (
	"net/http"
	"net/url"
)

type dispatcher struct {
	f func(http.ResponseWriter, *http.Request)
}

func (d *dispatcher) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	d.f(rw, r)
}

type Route struct {
	Path    string
	Handler http.Handler
}

func NewRoute(p string, f func(http.ResponseWriter, *http.Request)) *Route {
	return &Route{
		p,
		&dispatcher{f},
	}
}

func (r *Route) handles(uri *url.URL) bool {
	// Do matching things
	return true
}
