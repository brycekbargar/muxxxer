// Package muxxxer is a naive mux implementation with a single default mux
package muxxxer

import (
	"fmt"
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

// ServeHTTP dispatches requests to the registered Routes
func (m *muxxxer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	for _, r := range Default.routes {
		if r.handles(req.URL) {
			r.Handler.ServeHTTP(rw, req)
			return
		}
	}

	fmt.Fprintf(rw, "Muxxxer couldn't find a registration for %s", req.URL)
	// Do actual 404 things
}
