package muxxxer

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func matches(t *testing.T, r *Route, us string) {
	t.Helper()

	u, err := url.Parse(us)
	require.NoError(t, err)
	assert.True(t, r.handles(u), "because route: %s matches url: %s", r.Path, u.Path)
}

func doesnotmatch(t *testing.T, r *Route, us string) {
	t.Helper()

	u, err := url.Parse(us)
	require.NoError(t, err)
	assert.False(t, r.handles(u), "because route: %s does not match url: %s", r.Path, u.Path)
}

func TestRoute_exact(t *testing.T) {
	testCases := []struct {
		Name string
		Url  string
		Test func(*testing.T, *Route, string)
	}{
		{
			"Short Path",
			"http://localhost/some/path",
			matches,
		},
		{
			"Short Path with query parameters",
			"http://localhost/some/path?with=two&query=params",
			matches,
		},
		{
			"Short Path with trailing slash",
			"http://localhost/some/path/",
			matches,
		},
		{
			"Long Path",
			"http://localhost/some/path/that/goes/on",
			doesnotmatch,
		},
	}

	route, err := NewRoute("/some/path", nil)
	require.NoError(t, err)

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			tc.Test(t, route, tc.Url)
		})
	}
}

func TestRoute_trailingslash(t *testing.T) {
	testCases := []struct {
		Name string
		Url  string
		Test func(*testing.T, *Route, string)
	}{
		{
			"Short Path",
			"http://localhost/some/path",
			matches,
		},
		{
			"Short Path with query parameters",
			"http://localhost/some/path?with=two&query=params",
			matches,
		},
		{
			"Short Path with trailing slash",
			"http://localhost/some/path/",
			matches,
		},
		{
			"Long Path",
			"http://localhost/some/path/that/goes/on",
			matches,
		},
	}

	route, err := NewRoute("/some/path/", nil)
	require.NoError(t, err)

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			tc.Test(t, route, tc.Url)
		})
	}
}
