package muxxxer

import (
	"net/http"
	"net/url"
	"regexp"
)

var argRxp = regexp.MustCompile(`;[^\/]+(\/)?`)
var argMatcher = `[^\/]+$1`

// NewArgumentRoute creates a route with access to an ArgumentBag
//
// In order to capture a path segment to later be access in the ArgumentBag use the following syntax:
// "/some/path/;arg" where the semi-colon denotes the path segment as an argument
// Arguments can be located anywhere e.g. "/;object/path-stuff/;id/more-path-stuff/"
// The name of the captured Argument is everything in the segment after the semi-colon.
//
// Be careful! Semi-colons are things that are valid url characters,
// this matching scheme ignores reality in this respect for a simpler implementation
func NewArgumentRoute(r string, f func(http.ResponseWriter, *http.Request, *ArgumentBag)) (route *Route, err error) {
	route, err = NewRoute(
		argRxp.ReplaceAllString(r, argMatcher),
		func(rw http.ResponseWriter, req *http.Request) {
			f(rw, req, &ArgumentBag{rawRoute: r, url: req.URL})
		})
	if err != nil {
		return
	}

	route.rawPath = r
	return
}

// ArgumentBag contains ways to get query and path args from the matching url.
type ArgumentBag struct {
	rawRoute string
	url      *url.URL
	Args     map[string][]interface{}
}

// Parse populates Args.
//
// The resulting bag will contain the named parameters in the path in addition to any query parameters.
// These values are strings unless otherwise specified.
func (r *ArgumentBag) Parse() error {
	args := make(map[string][]interface{})

	for q, vs := range r.url.Query() {
		for _, v := range vs {
			args[q] = append(args[q], v)
		}
	}

	r.Args = args
	return nil
}
