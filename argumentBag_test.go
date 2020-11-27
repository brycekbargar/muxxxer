package muxxxer

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArgumentBag(t *testing.T) {
	testCases := []struct {
		Name           string
		Route          string
		Url            string
		ExpectedValues map[string][]string
	}{
		{
			"No arguments",
			"/a/path/",
			"http://localhost/a/path",
			make(map[string][]string),
		},
		{
			"Single query",
			"/a/path/",
			"http://localhost/a/path?pancakes=banana",
			map[string][]string{
				"pancakes": {"banana"},
			},
		},
		{
			"Multiple query",
			"/a/path/",
			"http://localhost/a/path?pancakes=banana&loris=slow",
			map[string][]string{
				"pancakes": {"banana"},
				"loris":    {"slow"},
			},
		},
		{
			"Same query value",
			"/a/path/",
			"http://localhost/a/path?pancakes=banana&pancakes=pecan",
			map[string][]string{
				"pancakes": {"banana", "pecan"},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			url, err := url.Parse(tc.Url)
			require.NoError(t, err)

			bag := &ArgumentBag{rawRoute: tc.Route, url: url}
			err = bag.Parse()
			require.NoError(t, err, "because the parsing should be ok")

			for n, vs := range tc.ExpectedValues {
				require.Contains(t, bag.Args, n)
				assert.ElementsMatch(t, vs, bag.Args[n])
			}
		})
	}

}
