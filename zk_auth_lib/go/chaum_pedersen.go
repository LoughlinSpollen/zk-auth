package chaum_pedersen

import (
	"crypto/rand"
	"fmt"
	"math/big"

	log "github.com/sirupsen/logrus"
)

type ChaumPedersenProver interface {
	Register(r *big.Int) (*big.Int, *big.Int)
	Response(s *big.Int) *big.Int
}

type ChaumPedersenVerifier interface {
	Challenge() *big.Int
	Verify(y1, y2, z *big.Int) bool
}

type chaumPedersen struct {
	g *big.Int
	q *big.Int
	a *big.Int
	b *big.Int
	A *big.Int
	B *big.Int
	C *big.Int
	r *big.Int
	s *big.Int
}

func NewChaumPedersen(g, q, a, b *big.Int) *chaumPedersen {

	A := new(big.Int).Exp(g, a, q)                      // g^a mod q
	B := new(big.Int).Exp(g, b, q)                      // g^b mod q
	C := new(big.Int).Exp(g, new(big.Int).Mul(a, b), q) // g^(a*b) mod q

	log.Trace(fmt.Printf("Public agree (g, g^a, g^b and g^ab) : %v, %v, %v, %v \n", g, A, B, C))
	return &chaumPedersen{
		g: g,
		q: q,
		a: a,
		b: b,
		A: A,
		B: B,
		C: C,
	}
}

func (cpd *chaumPedersen) Register(r *big.Int) (*big.Int, *big.Int) {
	cpd.r = r
	y1 := new(big.Int).Exp(cpd.g, cpd.r, cpd.q) // g^r mod q
	y2 := new(big.Int).Exp(cpd.B, cpd.r, cpd.q) // (g^b)^r mod q
	log.Trace(fmt.Printf("Prover computes and sends y1 y2 to Verifier (g^r, B^r)=(%v,%v)\n", y1, y2))
	return y1, y2
}

func (cpd *chaumPedersen) Response(s *big.Int) *big.Int {
	// // (r + a*s) mod q
	z := new(big.Int).Mod(cpd.r.Add(cpd.r, cpd.a.Mul(cpd.a, s)), cpd.q)
	log.Trace(fmt.Printf("Prover computes z=r+as (mod q)=%v\n", z))
	return z
}

func (cpd *chaumPedersen) Challenge() *big.Int {
	cpd.s, _ = rand.Int(rand.Reader, big.NewInt(1000))
	log.Trace(fmt.Printf("Verifier sends a challenge (s)=%v\n", cpd.s))
	return cpd.s
}

func (cpd *chaumPedersen) Verify(y1, y2, z *big.Int) bool {
	log.Trace(fmt.Printf("y1 %v ", y1))  // g^r mod q
	log.Trace(fmt.Printf("y2 %v\n", y2)) // (g^b)^r mod q
	log.Trace(fmt.Printf("Verifier makes two checks using the answer (z)=%v against the challenge (s)=%v\n", z, cpd.s))

	a1 := new(big.Int).Exp(cpd.g, z, cpd.q)
	log.Trace(fmt.Printf("Verifier calculates a1 g^z mod q=%d\n", a1))

	// A^s y1 mod q
	a2 := new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Exp(cpd.A, cpd.s, nil), y1), cpd.q)
	log.Trace(fmt.Printf("Verifier calculates a2: A^s y1=%d\n", a2))

	b1 := new(big.Int).Exp(cpd.B, z, cpd.q)
	log.Trace(fmt.Printf("Verifier calculates b1: B^z=%d\n", b1))

	// C^s y2 mod q
	b2 := new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Exp(cpd.C, cpd.s, nil), y2), cpd.q)
	log.Trace(fmt.Printf("Verifier calculates b2: C^s y2=%d\n", b2))

	return a1.Cmp(a2) == 0 && b1.Cmp(b2) == 0
}
