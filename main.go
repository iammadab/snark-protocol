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
	}
}

func (verifier *Verifier) Setup() []int64 {
	// evaluate the t(x) with unencrypted value of x
	verifier.EvalT = verifier.PolyT.EvaluateAt(verifier.EvalPoint)

	encryptedPowers := []int64{}
	for i := 0; i <= verifier.PolyT.Degree(); i++ {
		power := IntPow(verifier.EvalPoint, int64(i))
		encryptedPowers = append(encryptedPowers, verifier.EncryptValue(power))
	}

	fmt.Println("encrypted powers of x")
	fmt.Printf("%v", encryptedPowers)
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
	field := field.NewField(7)

	// need a more intiutive way to set the co-efficients
	// p(x) = x^3 - 3x^2 + 2x [0, 2, -3, 1]
	// h(x) = x^2 - 2x [0, -2, 1, 0]
	// t(x) = x - 1 [-1, 1, 0, 0]

	t_of_x := []int64{-1, 1, 0, 0}
	h_of_x := []int64{0, -2, 1, 0}
	p_of_x := []int64{0, 2, -3, 1}

	verifier := NewVerifier(field, 5, t_of_x)

	encrypted_powers_of_x := verifier.Setup()

	prover := NewProver(field, p_of_x, h_of_x)
	fmt.Println(prover)
	fmt.Println(prover.Prove(encrypted_powers_of_x))
	println(verifier.Verify(prover.Prove(encrypted_powers_of_x)))
}

// Helper function to perfom integer exponentiation in golang
func IntPow(a, b int64) int64 {
	return int64(math.Pow(float64(a), float64(b)))
}
