package web

import (
	"github.com/valyala/fasthttp"
	"encoding/json"
	"strconv"
	"context"
	"fmt"
)

//
type WebApp struct {

	Ctx context.Context

	//
	Response struct {
		Code int         `json:"code"`
		Body interface{} `json:"body"`
	}
}

//
func (a *WebApp) RenderJSONFastHttp(ctx *fasthttp.RequestCtx) error {
	ctx.Response.Header.SetCanonical([]byte(`Content-Type`), []byte(`application/json`))
	ctx.SetStatusCode(a.Response.Code)

	response, err := json.Marshal(a.Response)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody([]byte(`Internal Server Error`))
		return err
	}
	ctx.SetBody(response)
	return nil
}

//
func (a *WebApp) UserValueInt(ctx *fasthttp.RequestCtx, key string) (val int64) {
	valStr := fmt.Sprintf(`%s`, ctx.UserValue(key))
	val, err := strconv.ParseInt(valStr, 10,64)
	if err != nil {
		val = 0
	}
	return val
}

//
func (a *WebApp) UserValueUInt(ctx *fasthttp.RequestCtx, key string) (val uint64) {
	valStr := fmt.Sprintf(`%s`, ctx.UserValue(key))
	val, err := strconv.ParseUint(valStr, 10,64)
	if err != nil {
		val = 0
	}
	return val
}

//
func (a *WebApp) UserValueString(ctx *fasthttp.RequestCtx, key string) (string) {
	s :=  ctx.UserValue(key)
	if s == nil {
		return ""
	}
	return fmt.Sprintf(`%s`, s)
}