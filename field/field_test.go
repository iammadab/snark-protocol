package field

import "testing"

// testing the field, we want to make sure the correct answers are always gotten
// useful cases to test
// all the functions
// negative values, as this is possible.

type OperationTest struct {
	arg1, arg2, expected int
}

var AdditionTests = []OperationTest{
	{1, 2, 3},
	{-1, 2, 1},
	{-1, -2, 4},
	{-100, -200, 1},
}

func TestAddition(t *testing.T) {
	// look into the possibiblity of extracting the field into the test case
	// are there conditions where you'd want to test for different fields
	field := NewField(7)
	for _, test := range AdditionTests {
		if result := field.Add(int64(test.arg1), int64(test.arg2)); result != int64(test.expected) {
			t.Errorf("%d + %d = %d mod %d, expected %d", test.arg1, test.arg2, result, field.Order, test.expected)
		}
	}
}
