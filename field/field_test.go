package field

import "testing"

type OperationTest struct {
	arg1, arg2, expected int
}

var additionTests = []OperationTest{
	// {a, b, c} checks that a + b = c in the field
	{1, 2, 3},
	{-1, 2, 1},
	{-1, -2, 4},
	{-100, -200, 1},
}

var multiplicationTests = []OperationTest{
	// {a, b, c} checks that a * b = c in the field
	{92, 48, 6},
	{-1, 2, 5},
	{-46, 89, 1},
}

var exponentiationTests = []OperationTest{
	{2, 1, 2},
	{2, -1, 4},  // this is equivalent to finding the multiplicative inverse of 2
	{5, -10, 4}, // find the multiplicative inverse of 5 then raise to the power of 10
}

func TestAddition(t *testing.T) {
	field := NewField(7)
	for _, test := range additionTests {
		if result := field.Add(int64(test.arg1), int64(test.arg2)); result != int64(test.expected) {
			t.Errorf("%d + %d = %d mod %d, expected %d", test.arg1, test.arg2, result, field.Order, test.expected)
		}
	}
}

func TestMultiplication(t *testing.T) {
	field := NewField(7)
	for _, test := range multiplicationTests {
		if result := field.Mul(int64(test.arg1), int64(test.arg2)); result != int64(test.expected) {
			t.Errorf("%d * %d = %d mod %d, expected %d", test.arg1, test.arg2, result, field.Order, test.expected)
		}
	}
}

func TestExponentiation(t *testing.T) {
	field := NewField(7)
	for _, test := range exponentiationTests {
		if result := field.Exp(int64(test.arg1), int64(test.arg2)); result != int64(test.expected) {
			t.Errorf("%d ^ %d = %d mod %d, expected %d", test.arg1, test.arg2, result, field.Order, test.expected)
		}
	}
}
