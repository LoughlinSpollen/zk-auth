package chaum_pedersen_test

import (
	"math/big"
	"math/rand"
	"testing"

	cp "zk_auth_lib"
)

func TestChaumPedersen(t *testing.T) {
	q := big.NewInt(10003)
	g := big.NewInt(3)
	a := big.NewInt(10)
	b := big.NewInt(13)
	p := cp.NewChaumPedersen(g, q, a, b)
	v := cp.NewChaumPedersen(g, q, a, b)
	r := big.NewInt(rand.Int63n(1000))

	y1, y2 := p.Register(r)
	s := v.Challenge()
	z := p.Response(s)
	if !v.Verify(y1, y2, z) {
		t.Error("Failed to verify")
	}
}
