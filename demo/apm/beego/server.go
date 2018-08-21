package main

import (
	"github.com/astaxie/beego/context"
	"github.com/elastic/apm-agent-go"
	"github.com/elastic/apm-agent-go/module/apmhttp"
	"github.com/astaxie/beego"
)

type options struct {
	tracer         *elasticapm.Tracer
	requestIgnorer apmhttp.RequestIgnorerFunc
}

type middleware struct {
	tracer         *elasticapm.Tracer
	requestIgnorer apmhttp.RequestIgnorerFunc
}

func hello(ctx *context.Context) {
	ctx.WriteString("Hello!")
}

var UrlManager = func(c *context.Context) {

	opts := options{
		tracer:         elasticapm.DefaultTracer,
		requestIgnorer: apmhttp.DefaultServerRequestIgnorer(),
	}
	m := &middleware{
		tracer:         opts.tracer,
		requestIgnorer: opts.requestIgnorer,
	}
	req := c.Request
	name := req.Method + " " + req.RequestURI
	tx, req := apmhttp.StartTransaction(m.tracer, name, req)
	defer tx.End()
	tx.Result = apmhttp.StatusCodeResult(c.Output.Status)
	body := m.tracer.CaptureHTTPRequestBody(c.Request)

	if tx.Sampled() {
		tx.Context.SetHTTPRequest(c.Request)
		tx.Context.SetHTTPRequestBody(body)
		tx.Context.SetHTTPStatusCode(c.Output.Status)
		tx.Context.SetHTTPResponseHeaders(req.Header)
		tx.Context.SetHTTPResponseFinished(c.Output.IsOk())
		tx.Context.SetCustom("infoepoch", "beego")
	}
}

func main() {
	//beego.InsertFilter("/*", beego.BeforeRouter, UrlManager)
	beego.Any("/hello", hello)
	beego.Run(":8080")
}
