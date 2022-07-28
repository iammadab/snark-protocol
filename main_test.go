package main

import (
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

			encryptedPowersOfX, shiftedPowersOfX := verifier.Setup()

			prover := NewProver(field, test.pOfX, test.hOfX)
			p, shiftedP, h := prover.Prove(encryptedPowersOfX, shiftedPowersOfX)
			proofsValidity := verifier.Verify(p, shiftedP, h)

			if proofsValidity != test.isValid {
				t.Errorf("Test: %d, expected verifier to say %t, instead got %t", i, test.isValid, proofsValidity)
			}
		}
	}
}

func TestPolynomialRestriction(t *testing.T) {
	prime := int64(17707)
	field := field.NewField(prime)
	generator := 5

	IterationCount := 1000

	for j := 0; j < IterationCount; j++ {
		verifier := NewVerifier(field, int64(generator), testCases[0].tOfX)
		encryptedPowersOfX, shiftedPowersOfX := verifier.Setup()

		prover := NewProver(field, testCases[0].pOfX, testCases[0].hOfX)
		p, shiftedP, h := prover.Prove(encryptedPowersOfX, shiftedPowersOfX)

		// show that verifier accepts with correct inputs
		proofIsValid := verifier.Verify(p, shiftedP, h)
		if proofIsValid != true {
			// verifier didn't accept, this should not happen
			t.Errorf("Verifier failed to accept a valid proof, completeness is broken")
		}

		// TODO: possible that random element might be same as final value
		// set shiftedP to some arbitrary value
		shiftedP = field.RandomElement()
		proofIsValid = verifier.Verify(p, shiftedP, h)
		if proofIsValid == true {
			// verifier accepted a false proof, this should not happen
			t.Errorf("Verifier accept a false proof, soundness is broken")
		}
	}
}
