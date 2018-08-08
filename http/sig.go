package http

import (
	"github.com/valyala/fasthttp"
)

//
type SigError struct {
	Msg string
}

//
func (e*SigError) Error() string {
	return e.Msg
}

//
type Pair struct {
	key []byte
	value []byte
}


//
func MakeSigFastHttp(ctx *fasthttp.RequestCtx, secret string) ([]byte, error) {
	var pairs = make([]Pair, ctx.QueryArgs().Len())
	ctx.QueryArgs().VisitAll(func(key, value []byte) {
		pairs = append(pairs, Pair{key, value})
	})
	return makeSig(pairs, secret)
}

//
func makeSig(pairs []Pair, secret string) ([]byte, error) {
	return nil, &SigError{`Error make sig`}
}
