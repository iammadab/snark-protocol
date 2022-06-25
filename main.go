package main

import (
	"github.com/iammadab/snark-protocol/field"
	"github.com/iammadab/snark-protocol/polynomial"
)

type Verifier struct {
	PolyT polynomial.Polynomial
	Field field.Field
	EvalT int64
}

func NewVerifier(field field.Field, coefficients []int64) *Verifier {
	return &Verifier{
		PolyT: *polynomial.NewPolynomial(&field, coefficients),
		Field: field,
	}
}

// performs the unencrypted evaluation of t
// and returns the encrypted powers of s based on the degree of the polynomial
// func (verifier *Verifier) Setup() {
// 	// verifier.EvalT = verifier.PolyT.EvaluatePowers()
// }

func main() {
	a := field.NewField(7)
	println(a.Exp(2, 3))
}
