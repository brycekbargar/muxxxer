package muxxxer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArgumentRoute_matches(t *testing.T) {
	testCases := []struct {
		Name  string
		Route string
		Url   string
		Test  func(*testing.T, *Route, string)
	}{
		{
			"Exact Match",
			"/some/;name",
			"http://localhost/some/path",
			matches,
		},
		{
			"Exact Match with trailing slash",
			"/some/;name",
			"http://localhost/some/path/",
			matches,
		},
		{
			"No match before",
			"/some/;name/with/more",
			"http://localhost/a/path/with/more",
			doesnotmatch,
		},
		{
			"No match after",
			"/some/;name/with/more",
			"http://localhost/some/path/is/different",
			doesnotmatch,
		},
		{
			"Multiple matches",
			"/some/;name/and/;id",
			"http://localhost/some/path/and/such/",
			matches,
		},
		{
			"Trailing slash on route",
			"/some/;name/",
			"http://localhost/some/path/that/continues",
			matches,
		},
		{
			"No trailing slash on route",
			"/some/;name",
			"http://localhost/some/path/that/continues",
			doesnotmatch,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			route, err := NewArgumentRoute(tc.Route, nil)
			require.NoError(t, err)
			tc.Test(t, route, tc.Url)
		})
	}
}
