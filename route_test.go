package muxxxer

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func matches(t *testing.T, r Route, u *url.URL) {
	t.Helper()
	t.Logf("route: %s == url: %s", r.Path, u.Path)
	assert.True(t, r.handles(u))
}

func doesnotmatch(t *testing.T, r Route, u *url.URL) {
	t.Helper()
	t.Logf("route: %s != url: %s", r.Path, u.Path)
	assert.False(t, r.handles(u))
}

func TestRoute_handlesexact(t *testing.T) {
	route := Route{"/some/path", nil}

	url, _ := url.Parse("http://localhost/some/path")
	matches(t, route, url)

	url, _ = url.Parse("http://localhost/some/path/")
	matches(t, route, url)

	url, _ = url.Parse("http://localhost/some/path/that/goes/on")
	doesnotmatch(t, route, url)
}

func TestRoute_handletrailingslash(t *testing.T) {
	route := Route{"/some/path/", nil}

	url, _ := url.Parse("http://localhost/some/path")
	matches(t, route, url)

	url, _ = url.Parse("http://localhost/some/path/")
	matches(t, route, url)

	url, _ = url.Parse("http://localhost/some/path/that/goes/on")
	matches(t, route, url)
}
