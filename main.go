package main

import (
	"github.com/iammadab/snark-protocol/field"
	"github.com/iammadab/snark-protocol/polynomial"
)

type Verifier struct {
	PolyT     polynomial.Polynomial
	Field     *field.Field
	EvalPoint int64
	EvalT     int64
}

func NewVerifier(field *field.Field, coefficients []int64) *Verifier {
	return &Verifier{
		PolyT:     *polynomial.NewPolynomial(field, coefficients),
		Field:     field,
		EvalPoint: field.RandomElement(),
	}
}

// performs the unencrypted evaluation of t
// and returns the encrypted powers of s based on the degree of the polynomial
func (verifier *Verifier) Setup() {
	verifier.EvalT = verifier.PolyT.EvaluateAt(verifier.EvalPoint)
}

func main() {
	a := field.NewField(7)
	// println(a.Exp(2, 3))
	b := NewVerifier(a, []int64{5, 2})
	b.Setup()
	// fmt.Printf("%+v", b)
}
