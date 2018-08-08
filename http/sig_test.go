package http

import (
	"testing"
)


//
func TestMakeSigFastHttp(t *testing.T) {

}

//
func TestMakeSig(t *testing.T) {
	var (
		pairs = []Pair{
			Pair{[]byte(`xyz`) , []byte(`test`) },
			Pair{[]byte(`abc`) , []byte(`test2`) },
		}
		okSig = `xxx`
	)
	sig, err  := makeSig(pairs, `somesig`)
	if err != nil {
		t.Fatal(err)
	}

	if okSig != string(sig) {
		t.Fatal("okSig != string(sig)")
	}
}