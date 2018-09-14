package postmark

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/awnumar/memguard"
	"goji.io"
)

var (
	tMux    = goji.NewMux()
	tServer *httptest.Server
	client  *Client
)

func init() {
	tServer = httptest.NewServer(tMux)

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			// Reroute...
			return url.Parse(tServer.URL)
		},
	}

	testGuard, err := memguard.NewImmutableFromBytes([]byte("test"))
	if err != nil {
		panic(err.Error())
	}
	client = NewClient(testGuard, testGuard)
	client.HTTPClient = &http.Client{Transport: transport}
	client.BaseURL = tServer.URL
}
