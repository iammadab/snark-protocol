package main

import (
	"fmt"
	"math"

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
		//EvalPoint: 8,
	}
}

func (verifier *Verifier) Setup() []int64 {
	println("evaluation point", verifier.EvalPoint)

	// evaluate t(x) with unencrypted value of x
	verifier.EvalT = verifier.PolyT.EvaluateAt(verifier.EvalPoint)
	println("unencrypted evaluation of t(x)", verifier.EvalT)

	encryptedPowers := []int64{}
	for i := 0; i <= verifier.PolyT.Degree(); i++ {
		power := IntPow(verifier.EvalPoint, int64(i))
		println("powers of x", power)
		encryptedPowers = append(encryptedPowers, verifier.EncryptValue(power))
	}

	fmt.Printf("encrypted powers of x %v+\n", encryptedPowers)
	return encryptedPowers
}

func (verifier *Verifier) Verify(encryptedP, encryptedH int64) bool {
	println("p:", encryptedP)
	println("h:", encryptedH)
	println("t:", verifier.EvalT)
	println()

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

// Helper function to perfom integer exponentiation in golang
func IntPow(a, b int64) int64 {
	return int64(math.Pow(float64(a), float64(b)))
}
