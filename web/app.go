package web

import (
	"github.com/valyala/fasthttp"
	"encoding/json"
)

//
type WebApp struct {
	//
	Response struct {
		Code int         `json:"code"`
		Body interface{} `json:"body"`
	}
}

//
func (a *WebApp) RenderJSONFastHttp(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetCanonical([]byte(`Content-Type`), []byte(`application/json`))
	ctx.SetStatusCode(a.Response.Code)

	response, err := json.Marshal(a.Response)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody([]byte(`Internal Server Error`))
		return
	}
	ctx.SetBody(response)
}
