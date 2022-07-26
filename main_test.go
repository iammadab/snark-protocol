package main

import (
	"github.com/iammadab/snark-protocol/polynomial"
	"testing"

	"github.com/iammadab/snark-protocol/field"
)

type TestCase struct {
	t_of_x, p_of_x, h_of_x []int64
	is_valid               bool
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

func TestProtocol(t *testing.T) {
	// parameters to the functions seems to have a big effect, how do we know what to pick
	prime := int64(17707)
	field := field.NewField(prime)
	generator := 5

	ITERATION_COUNT := 1000

	for j := 0; j < ITERATION_COUNT; j++ {
		for i, test := range testCases {
			verifier := NewVerifier(field, int64(generator), test.t_of_x)

			encrypted_powers_of_x := verifier.Setup()

			prover := NewProver(field, test.p_of_x, test.h_of_x)
			p, h := prover.Prove(encrypted_powers_of_x)
			proofs_validity := verifier.Verify(p, h)

			if proofs_validity != test.is_valid {
				t.Errorf("Test: %d, expected verifier to say %t, instead got %t", i, test.is_valid, proofs_validity)
			}
		}
	}
}

func TestBreakHE(t *testing.T) {
	prime := int64(17707)
	field := field.NewField(prime)
	generator := 5

	ITERATION_COUNT := 1000

	for j := 0; j < ITERATION_COUNT; j++ {
		verifier := NewVerifier(field, int64(generator), testCases[0].t_of_x)
		encryptedPowersOfX := verifier.Setup()

		// Generate fake proof that fools the verifier with 100% probability
		randomPoint := field.RandomElement()
		encryptedH := EncryptValue(randomPoint, int64(generator), field)
		PolyT := polynomial.NewPolynomial(field, testCases[0].t_of_x)
		encryptedT := PolyT.EvaluateEncryptedPowers(encryptedPowersOfX)
		encryptedP := field.Exp(encryptedT, randomPoint)

		trickedVerifier := verifier.Verify(encryptedP, encryptedH)
		if trickedVerifier != true {
			t.Errorf("Failed to convince the verifier of a false proof")
		}
	}
}
