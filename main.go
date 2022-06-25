package main

type Field struct {
	Order int64
}

func NewField(order int64) *Field {
	return &Field{
		Order: order,
	}
}

func (field *Field) Add(a, b int64) int64 {
	return (a + b) % field.Order
}

func (field *Field) Sub(a, b int64) int64 {
	return (a - b) % field.Order
}

func (field *Field) Mul(a, b int64) int64 {
	return (a * b) % field.Order
}

func (field *Field) Div(a, b int64) int64 {
	return (a * field.MultiplicativeInverse(b)) % field.Order
}

func (field *Field) Exp(a, pow int64) int64 {
	// using repeated multiplication, more space efficient
	// TODO: is there something better
	result := int64(1)
	for i := int64(0); i < pow; i++ {
		result = (result * a) % field.Order
	}
	return result
}

// TODO: make the logic here clearer
func (field *Field) MultiplicativeInverse(b int64) int64 {
	a := field.Order
	if a < b {
		a, b = b, a
	}
	sa := [...]int64{1, 0}
	ta := [...]int64{0, 1}
	for b != 0 {
		q := a / b
		a, b = b, a%b
		sa[0], sa[1] = sa[1], sa[0]-q*sa[1]
		ta[0], ta[1] = ta[1], ta[0]-q*ta[1]
	}
	return ta[0]
}

func main() {
	a := NewField(7)
	println(a.Exp(2, 3))
}
