package http

import (
	"testing"
	"strings"
	"github.com/valyala/fasthttp"
)

const (
	T_OK_SIG = `d681883660b676605559901b7c99f4ea`
	T_SECRET = `somesig`
)

//
func TestMakeSigFastHttp(t *testing.T) {
	cont := fasthttp.RequestCtx{}
	cont.QueryArgs().Set(`xyz`,`test`)
	cont.QueryArgs().Set(`abc`,`test2`)
	cont.QueryArgs().Set(`sig`,T_OK_SIG)

	if !CheckSigFastHttp(&cont, T_SECRET, `sig`) {
		t.Error("!CheckSigFastHttp(&cont, T_SECRET, `sig`)")
	}

	cont.QueryArgs().Del(`sig`)
	cont.QueryArgs().Set(`sig`, `xxxxxxxxxxxx`)

	if CheckSigFastHttp(&cont, T_SECRET, `sig`) {
		t.Error("CheckSigFastHttp(&cont, T_SECRET, `sig`)")
	}
}

//
func TestMakeSig(t *testing.T) {

	var (
		pairs = []Pair{
			Pair{[]byte(`xyz`) , []byte(`test`) },
			Pair{[]byte(`abc`) , []byte(`test2`) },
		}
	)
	sig, err  := makeSig(pairs, T_SECRET)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.EqualFold(sig, T_OK_SIG) {
		t.Errorf("okSig(%s) != tSig(%s)", T_OK_SIG, sig )
	}

	if strings.EqualFold(sig, `xxxxxxxxxxxxxx`) {
		t.Errorf("strings.EqualFold(%s, `xxxxxxxxxxxxxx`", sig )
	}
}