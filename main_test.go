package main

import (
	"github.com/iammadab/snark-protocol/polynomial"
	"testing"

	"github.com/iammadab/snark-protocol/field"
)

type TestCase struct {
	tOfX, pOfX, hOfX []int64
	isValid          bool
}

var testCases = []TestCase{
	// TODO: write a simple compiler to accept polynomial in written form
	// t(x) = x - 1 [-1, 1, 0, 0]
	// p(x) = x^3 - 3x^2 + 2x [0, 2, -3, 1]
	// h(x) = x^2 - 2x [0, -2, 1, 0]
	{
		[]int64{-1, 1, 0, 0},
		[]int64{0, 2, -3, 1},
		[]int64{0, -2, 1, 0},
		true,
	},

	// t(x) = x - 1 [-1, 1, 0, 0]
	// p(x) = x^3 - 3x^2 + 3x [0, 3, -3, 1]
	// h(x) = x^2 - 3x [0, -3, 1, 0]
	{
		[]int64{-1, 1, 0, 0},
		[]int64{0, 3, -3, 1},
		[]int64{0, -3, 1, 0},
		false,
	},
}

// TODO: Completeness property breaks when I use larger primes e.g. 210403
//			only soundness broke before
func TestProtocol(t *testing.T) {
	// parameters to the functions seems to have a big effect, how do we know what to pick
	prime := int64(17707)
	field := field.NewField(prime)
	generator := 5

	IterationCount := 1000

	for j := 0; j < IterationCount; j++ {
		for i, test := range testCases {
			verifier := NewVerifier(field, int64(generator), test.tOfX)

			encryptedPowersOfX := verifier.Setup()

			prover := NewProver(field, test.pOfX, test.hOfX)
			p, h := prover.Prove(encryptedPowersOfX)
			proofsValidity := verifier.Verify(p, h)

			if proofsValidity != test.isValid {
				t.Errorf("Test: %d, expected verifier to say %t, instead got %t", i, test.isValid, proofsValidity)
			}
		}
	}
}

// TODO: Find the relationship between the field and 100% probability here
// 		changed to f* and is passing all the time.
func TestBreakHE(t *testing.T) {
	prime := int64(17707)
	field := field.NewField(prime)
	generator := 5

	IterationCount := 1000

	for j := 0; j < IterationCount; j++ {
		verifier := NewVerifier(field, int64(generator), testCases[0].tOfX)
		encryptedPowersOfX := verifier.Setup()

		// Generate fake proof that fools the verifier with 100% probability
		randomPoint := field.RandomElement()
		encryptedH := EncryptValue(randomPoint, int64(generator), field)
		PolyT := polynomial.NewPolynomial(field, testCases[0].tOfX)
		encryptedT := PolyT.EvaluateEncryptedPowers(encryptedPowersOfX)
		encryptedP := field.Exp(encryptedT, randomPoint)

		trickedVerifier := verifier.Verify(encryptedP, encryptedH)
		if trickedVerifier != true {
			// Print parameters in case of failure
			println("r", randomPoint)
			println("g^r", encryptedH)
			println("g^t", encryptedT)
			println("g^t^r", encryptedP)
			println("unencrypted t", PolyT.EvaluateAt(verifier.EvalT))
			t.Errorf("Failed to convince the verifier of a false proof")
		}
	}
}
