package muxxxer

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type dispatcher struct {
	f func(http.ResponseWriter, *http.Request)
}

func (d *dispatcher) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	d.f(rw, r)
}

// Route is a Regular Expression for testing whether the current request's Url matches the Route.
// When a route is matches the Handler is executed.
type Route struct {
	rawPath string
	Path    *regexp.Regexp
	Handler http.Handler
}

// NewRoute creates a new Route by converting the handlerFunc to a Handler
// Errors returned are from the regexp.Regexp library
func NewRoute(p string, f func(http.ResponseWriter, *http.Request)) (route *Route, err error) {
	route = &Route{rawPath: p, Handler: &dispatcher{f}}

	p = strings.TrimLeft(p, "/")
	p = `\/?` + p

	if strings.HasSuffix(p, "/") {
		p = strings.TrimRight(p, "/")
		p += `(\/.*)?`
	} else {
		p += `\/?`
	}

	route.Path, err = regexp.Compile("^" + p + "$")

	return
}

func (r *Route) handles(uri *url.URL) bool {
	return r.Path.MatchString(uri.Path)
}
