package field

import "testing"

// testing the field, we want to make sure the correct answers are always gotten
// useful cases to test
// all the functions
// negative values, as this is possible.

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

// TODO: you don't need to duplicate this function
// abstract everything into testcase (might be difficult to read tho)
func TestAddition(t *testing.T) {
	// look into the possibiblity of extracting the field into the test case
	// are there conditions where you'd want to test for different fields
	field := NewField(7)
	for _, test := range additionTests {
		if result := field.Add(int64(test.arg1), int64(test.arg2)); result != int64(test.expected) {
			t.Errorf("%d + %d = %d mod %d, expected %d", test.arg1, test.arg2, result, field.Order, test.expected)
		}
	}
}

func TestMultiplication(t *testing.T) {
	// look into the possibiblity of extracting the field into the test case
	// are there conditions where you'd want to test for different fields
	field := NewField(7)
	for _, test := range multiplicationTests {
		if result := field.Mul(int64(test.arg1), int64(test.arg2)); result != int64(test.expected) {
			t.Errorf("%d + %d = %d mod %d, expected %d", test.arg1, test.arg2, result, field.Order, test.expected)
		}
	}
}
