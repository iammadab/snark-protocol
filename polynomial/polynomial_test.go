package polynomial

import (
	"testing"

	"github.com/iammadab/snark-protocol/field"
)

type PolynomialTest struct {
	poly     *Polynomial
	point    int64
	expected int64
}

func TestPolynomialEvaluationAtAPoint(t *testing.T) {
	f := field.NewField(7)

	// f(x) = x
	// constant: 0, co(x) = 1
	poly := NewPolynomial(f, []int64{0, 1})
	cases := []PolynomialTest{
		{poly, 1, 1},
		{poly, 7, 0},
		{poly, 8, 1},
	}

	// f(x) = 2x + 1
	// constant = 1 co(x) = 2
	poly = NewPolynomial(f, []int64{1, 2})
	cases = append(cases, []PolynomialTest{
		{poly, 1, 3},
		{poly, 7, 1},
		{poly, 8, 3},
	}...)

	for _, test := range cases {
		if result := test.poly.EvaluateAt(test.point); result != test.expected {
			t.Errorf("Evaluattion of %+v at %d mod %d = %d, expected %d", test.poly.Coefficients, test.point, test.poly.Field.Order, result, test.expected)
		}

	}

	// polynomial := NewPolynomial(f, []int64{0, 1})
	// if polynomial.EvaluateAt(1) != 1 {
	// 	t.Errorf("evaluation failed")
	// }
}
