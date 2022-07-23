package main

import (
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

func TestProtocol(t *testing.T) {
	field := field.NewField(7919)
	generator := 5

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
