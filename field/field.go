package field

import (
	"math/rand"
	"time"
)

type Field struct {
	Order int64
}

func NewField(order int64) *Field {
	return &Field{
		Order: order,
	}
}

func (field *Field) Mod(val int64) int64 {
	element := val % field.Order
	if element < 0 {
		return element + field.Order
	}
	return element
}

func (field *Field) Add(a, b int64) int64 {
	return field.Mod(field.Mod(a) + field.Mod(b))
}

func (field *Field) Sub(a, b int64) int64 {
	return field.Mod(field.Mod(a) - field.Mod(b))
}

func (field *Field) Mul(a, b int64) int64 {
	return field.Mod(field.Mod(a) * field.Mod(b))
}

func (field *Field) Exp(a, pow int64) int64 {
	// convert negative exponent problem instances to their positive versions
	// by finding the multiplicative inverse of the base, and converting the exponent
	// to its positive version.
	// a^-b => (a^-1)^b
	// a^-1 is the multiplicative inverse of a
	if pow < 0 {
		a = field.MultiplicativeInverse(a)
		pow *= -1
	}

	return field.FastExp(a, pow)

}

func (field *Field) FastExp(base, pow int64) int64 {
	// base^0 = 1
	if pow == 0 {
		return 1
	}

	// base^1 = base
	if pow == 1 {
		return field.Mod(base)
	}

	temp := field.FastExp(base, pow/2)
	result := field.Mul(temp, temp)

	if pow%2 == 1 {
		result = field.Mul(result, base)
	}

	return result
}

// TODO: make this logic here clearer
// also, multiplicative inverses might not exist depending on the field
// return an error in such case
func (field *Field) MultiplicativeInverse(b int64) int64 {
	a := field.Order
	if a < b {
		a, b = b, a
	}
	ta := [...]int64{0, 1}
	for b != 0 {
		q := a / b
		a, b = b, a%b
		ta[0], ta[1] = ta[1], ta[0]-q*ta[1]
	}
	return ta[0]
}

// TODO: I changed this to f*, is this right??
func (field *Field) RandomElement() int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(field.Order-1) + 1
}
