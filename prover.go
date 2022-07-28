package main

import (
	"github.com/iammadab/snark-protocol/field"
	"github.com/iammadab/snark-protocol/polynomial"
)

type Prover struct {
	field *field.Field
	PolyP polynomial.Polynomial
	PolyH polynomial.Polynomial
}

func NewProver(field *field.Field, polyp []int64, polyh []int64) *Prover {
	return &Prover{
		field: field,
		PolyP: *polynomial.NewPolynomial(field, polyp),
		PolyH: *polynomial.NewPolynomial(field, polyh),
	}
}

func (prover *Prover) Prove(powers []int64, shiftedPowers []int64) (int64, int64, int64) {
	encryptedEvaluationOfP := prover.PolyP.EvaluateEncryptedPowers(powers)
	shiftedEvaluationOfP := prover.PolyP.EvaluateEncryptedPowers(shiftedPowers)
	encryptedEvaluationOfH := prover.PolyH.EvaluateEncryptedPowers(powers)

	// add zero knowledge
	blindingFactor := prover.field.RandomElement()
	encryptedEvaluationOfP = prover.field.Exp(encryptedEvaluationOfP, blindingFactor)
	shiftedEvaluationOfP = prover.field.Exp(shiftedEvaluationOfP, blindingFactor)
	encryptedEvaluationOfH = prover.field.Exp(encryptedEvaluationOfH, blindingFactor)

	return encryptedEvaluationOfP, shiftedEvaluationOfP, encryptedEvaluationOfH
}
