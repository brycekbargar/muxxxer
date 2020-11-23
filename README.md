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
	muxxxer.RegisterRoute(muxxxer.NewRoute("/index", Index))
	muxxxer.RegisterRoute(muxxxer.NewRoute(...))
	muxxxer.RegisterRoute(muxxxer.NewRoute(...))

	http.ListenAndServe(":8600", &muxxxer.Default)
}
```

### Path Matching

Like the go standard library the path has to be an exact match to the route, unless there is a trailing slash.
When a trailing slash is present in the route than any path beginning with that path is matched.