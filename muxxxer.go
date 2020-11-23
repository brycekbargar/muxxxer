package muxxxer

import (
	"fmt"
	"net/http"
)

type muxxxer struct {
	routes []*Route
}

var Default muxxxer

func init() {
	Default = muxxxer{
		make([]*Route, 0),
	}
}

func RegisterRoute(r *Route) {
	Default.routes = append(Default.routes, r)
}

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
