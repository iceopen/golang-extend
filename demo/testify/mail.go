package testify

import (
	"fmt"
	"net"
	"net/smtp"
	"strings"
)

type dialer interface {
	Close() error
	Hello(localName string) error
	Mail(from string) error
	Rcpt(to string) error
}

var (
	netLookupMX = net.LookupMX
	smtpClient  = func(addr string) (dialer, error) {
		// Dial the tcp connection
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			return nil, err
		}

		// Connect to the SMTP server
		c, err := smtp.NewClient(conn, addr)
		if err != nil {
			return nil, err
		}

		return c, nil
	}
)

func ValidateHost(email string) (err error) {
	mx, err := netLookupMX(host(email))
	if err != nil {
		return err
	}

	client, err := smtpClient(fmt.Sprintf("%s:%d", mx[0].Host, 25))
	if err != nil {
		return err
	}

	defer func() {
		if er := client.Close(); er != nil {
			err = er
		}
	}()

	if err = client.Hello("checkmail.me"); err != nil {
		return err
	}
	if err = client.Mail("testing-email-host@gmail.com"); err != nil {
		return err
	}
	return client.Rcpt(email)
}

func host(email string) (host string) {
	i := strings.LastIndexByte(email, '@')
	return email[i+1:]
}
