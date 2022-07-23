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

// TODO: maybe also perform modular operation on input before computation

func (field *Field) Add(a, b int64) int64 {
	return field.Mod(field.Mod(a) + field.Mod(b))
}

func (field *Field) Sub(a, b int64) int64 {
	return field.Mod(field.Mod(a) - field.Mod(b))
}

func (field *Field) Mul(a, b int64) int64 {
	return field.Mod(field.Mod(a) * field.Mod(b))
}

// TODO: re-enable
// func (field *Field) Div(a, b int64) int64 {
// 	return field.Mod(a * field.MultiplicativeInverse(b))
// }

func (field *Field) Exp(a, pow int64) int64 {
	// using repeated multiplication, more space efficient
	// TODO: is there something better
	if pow < 0 {
		a = field.MultiplicativeInverse(a)
		pow *= -1
	}

	return field.FastExp(a, pow)

}

func (field *Field) FastExp(a, pow int64) int64 {
	if pow == 0 {
		return 1
	}

	if pow == 1 {
		return field.Mod(a)
	}

	temp := field.FastExp(a, pow/2)
	result := field.Mul(temp, temp)

	if pow%2 == 1 {
		result = field.Mul(result, a)
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
	// sa := [...]int64{1, 0}
	ta := [...]int64{0, 1}
	for b != 0 {
		q := a / b
		a, b = b, a%b
		// sa[0], sa[1] = sa[1], sa[0]-q*sa[1]
		ta[0], ta[1] = ta[1], ta[0]-q*ta[1]
	}
	return ta[0]
}

func (field *Field) RandomElement() int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(field.Order)
}
