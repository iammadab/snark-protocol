package main

import (
	"github.com/iammadab/snark-protocol/field"
	"github.com/iammadab/snark-protocol/polynomial"
)

// TODO: need the concept of public information and private information
// 		generator, field, polyT should all be public

type Verifier struct {
	Generator  int64
	PolyT      polynomial.Polynomial
	Field      *field.Field
	EvalPoint  int64
	EvalT      int64
	ShiftValue int64
}

func NewVerifier(field *field.Field, generator int64, coefficients []int64) *Verifier {
	return &Verifier{
		Generator:  generator,
		PolyT:      *polynomial.NewPolynomial(field, coefficients),
		Field:      field,
		EvalPoint:  field.RandomElement(),
		ShiftValue: field.RandomElement(),
	}
}

// Setup allows the verifier to get the unencrypted evaluation of t_of_x
// and also generate the encrypted powers of s for the prover
func (verifier *Verifier) Setup() (encryptedPowers, shiftedEncryptedPowers []int64) {
	verifier.EvalT = verifier.PolyT.EvaluateAt(verifier.EvalPoint)
	encryptedPowers = GenerateEncryptedPowers(verifier.EvalPoint, verifier.PolyT.Degree(), verifier.Generator, verifier.Field)
	shiftedEncryptedPowers = ShiftEncryptedPowers(encryptedPowers, verifier.ShiftValue, verifier.Field)
	return
}

// Verify checks that polynomial restriction constraint and that p = ht in encrypted space
func (verifier *Verifier) Verify(encryptedP, shiftedP, encryptedH int64) bool {
	// check polynomial restriction constraint
	expectedShiftEval := verifier.Field.Exp(encryptedP, verifier.ShiftValue)
	if expectedShiftEval != shiftedP {
		// polynomial restriction check failed
		return false
	}

	// Check p = ht in encrypted space
	p := verifier.Field.Mod(encryptedP)
	ht := verifier.Field.Mod(verifier.Field.Exp(encryptedH, verifier.EvalT))

	return p == ht
}
