package testify

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

type smtpDialerMock struct {
}

func (*smtpDialerMock) Close() error {
	return nil
}

func (*smtpDialerMock) Hello(localName string) error {
	return nil
}

func (*smtpDialerMock) Mail(from string) error {
	return nil
}

func (*smtpDialerMock) Rcpt(to string) error {
	return nil
}

func TestValidateHost(t *testing.T) {
	netLookupMX = func(name string) ([]*net.MX, error) {
		mxs := []*net.MX{
			{
				Host: "host.tld",
				Pref: 1,
			},
		}

		return mxs, nil
	}

	smtpClient = func(addr string) (dialer, error) {
		client := &smtpDialerMock{}
		return client, nil
	}

	email := "mail@host.tld"
	actual := ValidateHost(email)
	assert.NoError(t, actual)
}

func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

}
