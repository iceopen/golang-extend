package examples

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/pingcap/fn"
)

type Request struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct{
	Token string `json:"token"`
}

func api1() (*Response, error) {
	return &Response{ Token: "token" }, nil
}

func api2(request *Request) (*Response, error) {
	token := request.Username + request.Password
	return &Response{ Token: token }, nil
}

func api3(rawreq *http.Request, request *Request) (*Response, error) {
	token := request.Username + request.Password
	return &Response{ Token: token }, nil
}

func api4(rawreq http.Header, request *Request) (*Response, error) {
	token := request.Username + request.Password
	return &Response{ Token: token }, nil
}

func api5(form *fn.Form, request *Request) (*Response, error) {
	token := request.Username + request.Password + form.Get("type")
	return &Response{ Token: token }, nil
}

func api6(body io.ReadCloser, request *Request) (*Response, error) {
	token := request.Username + request.Password
	return &Response{ Token: token }, nil
}

func api7(form *multipart.Form, request *Request) (*Response, error) {
	token := request.Username + request.Password
	return &Response{ Token: token }, nil
}

func api7(urls *url.URL, request *Request) (*Response, error) {
	token := request.Username + request.Password
	return &Response{ Token: token }, nil
}

func api8(urls *url.URL, form *multipart.Form, body io.ReadCloser, rawreq http.Header, request *Request) (*Response, error) {
	token := request.Username + request.Password
	return &Response{ Token: token }, nil
}