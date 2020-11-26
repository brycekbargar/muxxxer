package muxxxer

import (
	"net/http"
	"net/url"
	"regexp"
)

var argRxp = regexp.MustCompile(`;[^\/]+(\/)?`)
var argMatcher = `[^\/]+$1`

// ArgumentBag contains ways to get path and query args from the matching url.
type ArgumentBag struct {
	rawRoute string
	url      url.URL
}

func newArgumentBag(r string, req *http.Request) *ArgumentBag {
	return &ArgumentBag{r, *req.URL}
}

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
			f(rw, req, newArgumentBag(r, req))
		})
	if err != nil {
		return
	}

	route.rawPath = r
	return
}
