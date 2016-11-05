package httpcli

import (
	"net/url"
	"testing"

	"code.uber.internal/infra/mesos-go/encoding"
	"github.com/stretchr/testify/assert"
)

func Test_HTTPCli(t *testing.T) {
	// Setup api url
	apiURL := url.URL{
		Scheme: "http",
		Host:   "someip:0",
		Path:   "/some/endpint",
	}

	cli := New(
		Endpoint(apiURL.String()),
		Method("GET"),
		Codec(&encoding.JSONCodec),
		Do(With(Timeout(120))),
	)

	assert.Equal(t, cli.url, "http://someip:0/some/endpint")
	assert.Equal(t, cli.method, "GET")
}
