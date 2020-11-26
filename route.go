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

// Route is a combination of an http.Handler and the path it is registered to handle
type Route struct {
	Path    *regexp.Regexp
	Handler http.Handler
}

// NewRoute creates a new Route by converting the handlerFunc to a Handler
func NewRoute(p string, f func(http.ResponseWriter, *http.Request)) *Route {
	p = strings.TrimLeft(p, "/")

	if strings.HasSuffix(p, "/") {
		p = strings.TrimRight(p, "/")
		p += `(\/.*)?`
	} else {
		p += `\/?`
	}
	rxp, err := regexp.Compile("^" + p + "$")
	if err != nil {
		panic(err)
	}

	return &Route{
		rxp,
		&dispatcher{f},
	}
}

func (r *Route) handles(uri *url.URL) bool {
	return r.Path.MatchString(strings.TrimLeft(uri.Path, "/"))
}
