# Muxxxer

While working through "Go Web Programming" I wanted to see if I was understanding how the net/http library was working under the hood. Go is a little less magic than other web frameworks but is still pretty magic.

This mux "implementation" converts handlerFuncs into actual handlers using a dispatcher, I don't really know how the actual standard library does it but this seemed reasonable to me?

There's also a Default instance of the muxxer exposed. This default instance is what gets routes registered. This kind of static global pattern is new and a little weird to me but I wanted to see if I could duplicate the standard libraries behavior.

The Route struct is used to "encapsulate" the logic for matching a route and give a nice struct to pass around for path/handler pairs.

## Usage

I also wanted to test creating and linking packages with go, the server implementation using this mux package looks something like:

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/brycekbargar/muxxxer"
)

func Index(rw http.ResponseWriter, req *http.Request) {
    ...
}

func main() {
	muxxxer.MustRegisterRoute(muxxxer.NewRoute("/index", Index))
	muxxxer.MustRegisterRoute(muxxxer.NewRoute(...))
	muxxxer.MustRegisterRoute(muxxxer.NewRoute(...))

	http.ListenAndServe(":8600", &muxxxer.Default)
}
```

### Path Matching

Like the go standard library the path has to be an exact match to the route, unless there is a trailing slash.
When a trailing slash is present in the route than any path beginning with that path is matched.

If you want to specify your own regexp.Regexp for the Path note the url.Url.Path it is compared to is not modified in any way.

### Path Arguments

One thing I noticed was missing from the standard library was parameters from the path of a route.
For example, `/users/{id}` or `/users/{id}/actions/disable` etc.
I'm pretty sure actual go web frameworks have this and I wanted to see what implementing it would entail.

You can optionally register a route using the `muxxxer.NewArgumentRoute` function.
This modifies the normal ServeHTTP callback to also take an `ArgumentBag` parameter.
This bag requires Parsing like parsing of form values in the standard library.
After parsing, values captured from the route and query parameters are available in a map by name.

The map is a `map[string][]interface{}`. This is a little weird.
I had grand intentions to parse different kinds of values like integers and booleans which is why I chose a map of interfaces instead of strings.
