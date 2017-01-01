package httpcli

import (
	"io/ioutil"
	"net/url"
	"strings"
	"testing"

	"github.com/mesos/mesos-go/encoding"

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
		Codec(&encoding.JSONCodec),
		Do(With(Timeout(120))),
	)

	assert.Equal(t, cli.url, "http://someip:0/some/endpint")
}

type exampleStruct struct {
	StringField string
	IntField    int
}

func Test_IdentityReader(t *testing.T) {
	const bufSize = 32
	buf := make([]byte, bufSize)

	RandStringRunesInit()
	testString := RandStringRunes(bufSize + 1)

	// Setup the reader
	reader := IdentityReader{r: ioutil.NopCloser(strings.NewReader(testString))}

	// First read should not return eof, with n=bufSize
	eof, n, err := reader.ReadFrame(buf)
	assert.Nil(t, err)
	assert.False(t, eof)
	assert.Equal(t, bufSize, n)
	assert.Equal(t, testString[:bufSize], string(buf[:n]))

	// Second read should return eof, with n=1
	eof, n, err = reader.ReadFrame(buf)
	assert.Nil(t, err)
	assert.False(t, eof)
	assert.Equal(t, 1, n)
	assert.Equal(t, testString[bufSize:], string(buf[:n]))

	// Third read should return eof, with n=0
	eof, n, err = reader.ReadFrame(buf)
	assert.Nil(t, err)
	assert.True(t, eof)
	assert.Equal(t, 0, n)
}
