package main

import (
	"github.com/iammadab/snark-protocol/field"
	"github.com/iammadab/snark-protocol/polynomial"
	"math"
)

// TODO: need the concept of public information and private information
// 		generator, field, polyT should all be public

type Verifier struct {
	Generator int64
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

// Setup allows the verifier to get the unencrypted evaluation of t_of_x
// and also generate the encrypted powers of s for the prover
func (verifier *Verifier) Setup() []int64 {
	verifier.EvalT = verifier.PolyT.EvaluateAt(verifier.EvalPoint)
	return GenerateEncryptedPowers(verifier.EvalPoint, verifier.PolyT.Degree(), verifier.Generator, verifier.Field)
}

// Verify checks that p = ht in encrypted space
// E(p) = E(h)^t
func (verifier *Verifier) Verify(encryptedP, encryptedH int64) bool {
	p := verifier.Field.Mod(encryptedP)
	ht := verifier.Field.Mod(verifier.Field.Exp(encryptedH, verifier.EvalT))

	return p == ht
}

func (verifier *Verifier) EncryptValue(val int64) int64 {
	return verifier.Field.Exp(verifier.Generator, val)
}

// IntPow performs integer exponentiation
func IntPow(a, b int64) int64 {
	return int64(math.Pow(float64(a), float64(b)))
}
