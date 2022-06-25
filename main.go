package main

import (
	"fmt"

	"github.com/iammadab/snark-protocol/field"
	"github.com/iammadab/snark-protocol/polynomial"
)

type Verifier struct {
	Generator int64 // both prover and verifier should have this
	PolyT     polynomial.Polynomial
	Field     *field.Field
	EvalPoint int64
	EvalT     int64
}

func NewVerifier(field *field.Field, generator int64, coefficients []int64) *Verifier {
	return &Verifier{
		Generator: generator,
		PolyT:     *polynomial.NewPolynomial(field, coefficients),
		Field:     field,
		EvalPoint: field.RandomElement(),
	}
}

// performs the unencrypted evaluation of t
// and returns the encrypted powers of s based on the degree of the polynomial
func (verifier *Verifier) Setup() []int64 {
	verifier.EvalT = verifier.PolyT.EvaluateAt(verifier.EvalPoint)
	encryptedPowers := []int64{1}
	for i := 0; i < verifier.PolyT.Degree(); i++ {
		encryptedPowers = append(encryptedPowers, verifier.EncryptValue(verifier.EvalPoint))
	}
	return encryptedPowers
}

func (verifier *Verifier) Verify(encryptedP, encryptedH int64) bool {
	// TODO: create an equality function in field package
	return verifier.Field.Mod(encryptedP) == verifier.Field.Mod(verifier.Field.Exp(encryptedH, verifier.EvalT))
}

func (verifier *Verifier) EncryptValue(val int64) int64 {
	return verifier.Field.Exp(verifier.Generator, val)
}

type Prover struct {
	PolyP polynomial.Polynomial
	PolyH polynomial.Polynomial
}

func (prover *Prover) Prove(powers []int64) (int64, int64) {
	return prover.PolyP.EvaluateEncryptedPowers(powers), prover.PolyH.EvaluateEncryptedPowers(powers)
}

func NewProver(field *field.Field, polyp []int64, polyh []int64) *Prover {
	return &Prover{
		PolyP: *polynomial.NewPolynomial(field, polyp),
		PolyH: *polynomial.NewPolynomial(field, polyh),
	}
}

func main() {
	a := field.NewField(7)
	// println(a.Exp(2, 3))
	// 2x + 5
	b := NewVerifier(a, 5, []int64{5, 2})
	r := b.Setup()
	fmt.Println(r)
	// fmt.Printf("%+v", b)
	m := NewProver(a, []int64{1, 5}, []int64{5, 3})
	fmt.Println(m)
	fmt.Println(m.Prove(r))
	println(b.Verify(m.Prove(r)))
}
