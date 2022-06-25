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
	return field.Mod(a + b)
}

func (field *Field) Sub(a, b int64) int64 {
	return field.Mod(a - b)
}

func (field *Field) Mul(a, b int64) int64 {
	return field.Mod(a * b)
}

func (field *Field) Div(a, b int64) int64 {
	return field.Mod(a * field.MultiplicativeInverse(b))
}

func (field *Field) Exp(a, pow int64) int64 {
	// using repeated multiplication, more space efficient
	// TODO: is there something better
	result := int64(1)
	for i := int64(0); i < pow; i++ {
		result = field.Mod(result * a)
	}
	return result
}

// TODO: make the logic here clearer
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
