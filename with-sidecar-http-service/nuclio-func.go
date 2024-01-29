package main

import (
	nuclio "github.com/nuclio/nuclio-sdk-go"
	"github.com/valyala/fasthttp"
	"net/url"
)

func Handler(context *nuclio.Context, event nuclio.Event) (interface{}, error) {
	context.Logger.Info("Got request, sending it to sidecar container")
	req := fasthttp.AcquireRequest()
	req.SetBody(event.GetBody())
	sidecarHost, _ := url.JoinPath("http://0.0.0.0:9000", event.GetPath())
	context.Logger.Info(sidecarHost)
	req.SetRequestURI(sidecarHost)
	resp := fasthttp.AcquireResponse()
	err := fasthttp.Do(req, resp)
	return nuclio.Response{
		StatusCode:  resp.StatusCode(),
		ContentType: "application/text",
		Body:        resp.Body(),
	}, err
}
