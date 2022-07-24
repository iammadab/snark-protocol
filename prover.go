package main

import (
	"github.com/iammadab/snark-protocol/field"
	"github.com/iammadab/snark-protocol/polynomial"
)

type Prover struct {
	PolyP polynomial.Polynomial
	PolyH polynomial.Polynomial
}

func NewProver(field *field.Field, polyp []int64, polyh []int64) *Prover {
	return &Prover{
		PolyP: *polynomial.NewPolynomial(field, polyp),
		PolyH: *polynomial.NewPolynomial(field, polyh),
	}
}

func (prover *Prover) Prove(powers []int64) (int64, int64) {
	encryptedEvaluationOfP := prover.PolyP.EvaluateEncryptedPowers(powers)
	encryptedEvaluationOfH := prover.PolyH.EvaluateEncryptedPowers(powers)
	return encryptedEvaluationOfP, encryptedEvaluationOfH
}
