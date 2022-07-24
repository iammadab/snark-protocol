package main

import (
	"github.com/iammadab/snark-protocol/field"
	"github.com/iammadab/snark-protocol/polynomial"
	"math"
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

func (verifier *Verifier) Setup() []int64 {
	verifier.EvalT = verifier.PolyT.EvaluateAt(verifier.EvalPoint)

	encryptedPowers := []int64{}
	for i := 0; i <= verifier.PolyT.Degree(); i++ {
		power := IntPow(verifier.EvalPoint, int64(i))
		encryptedPowers = append(encryptedPowers, verifier.EncryptValue(power))
	}

	return encryptedPowers
}

func (verifier *Verifier) Verify(encryptedP, encryptedH int64) bool {
	return verifier.Field.Mod(encryptedP) == verifier.Field.Mod(verifier.Field.Exp(encryptedH, verifier.EvalT))
}

func (verifier *Verifier) EncryptValue(val int64) int64 {
	return verifier.Field.Exp(verifier.Generator, val)
}

// IntPow performs integer exponentiation
func IntPow(a, b int64) int64 {
	return int64(math.Pow(float64(a), float64(b)))
}
