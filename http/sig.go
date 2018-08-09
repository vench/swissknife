package http

import (
	"github.com/valyala/fasthttp"
	"strings"
	"bytes"
	"sort"
	"crypto/md5"
	"fmt"
)

const (
	SIG_GLUE_KEY_VALUE = '='
	SIG_GLUE_PAIR = '|'
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
func CheckSigFastHttp(ctx *fasthttp.RequestCtx, secret, sigKey string) bool {
	var pairs = make([]Pair, 0)
	ctx.QueryArgs().VisitAll(func(key, value []byte) {
		if !strings.EqualFold(sigKey, string(key)) {
			pairs = append(pairs, Pair{key, value})
		}
	})

	sig, err := makeSig(pairs, secret)
	if err != nil {
		return false
	}
	if !ctx.QueryArgs().Has(sigKey) {
		return false
	}

	if !strings.EqualFold(sig, string(ctx.QueryArgs().Peek(sigKey))) {
		return  false
	}

	return true
}

//
func makeSig(pairs []Pair, secret string) (string, error) {
	sort.Slice(pairs, func(i, j int) bool {
		return bytes.Compare(pairs[i].key, pairs[j].key) < 0
	})

	var (
		sig []byte
	)

	for n := 0; n < len(pairs); n ++ {
		if n > 0 {
			sig = append(sig, SIG_GLUE_PAIR)
		}
		sig = append(sig, pairs[n].key...)
		sig = append(sig, SIG_GLUE_KEY_VALUE)
		sig = append(sig, pairs[n].value...)
	}

	sig = append(sig, []byte(secret)...)
	h := md5.Sum(sig)
	return fmt.Sprintf(`%x`, h), nil
}
