// Package muxxxer is a naive mux implementation with a single default mux
//
// Note: This isn't intended for production use, it's just a test implementation for my own learning purposes
package muxxxer

import (
	"net/http"
)

type muxxxer struct {
	routes []*Route
}

// Default is the instance of the mux implementation to handle request
var Default muxxxer

func init() {
	Default = muxxxer{
		make([]*Route, 0),
	}
}

// RegisterRoute lets the Default mux know to route requests to the handler
func RegisterRoute(r *Route) {
	Default.routes = append(Default.routes, r)
}

// MustRegisterRoute lets the Default mux know to route requests to the handler
// It panics if given an error (this is probably an error during route instantiation)
func MustRegisterRoute(r *Route, err error) {
	if err != nil {
		panic(err)
	}

	Default.routes = append(Default.routes, r)
}

// ServeHTTP dispatches requests to the registered Routes
func (m *muxxxer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	for _, r := range Default.routes {
		if r.Path.MatchString(req.URL.Path) &&
			r.Handler != nil {
			r.Handler.ServeHTTP(rw, req)
			return
		}
	}

	http.NotFound(rw, req)
}
