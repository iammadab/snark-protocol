// TODO: add documentation
// all evaluations happen in mod space
package polynomial

import (
	"github.com/iammadab/snark-protocol/field"
)

type Polynomial struct {
	Field        *field.Field
	Coefficients []int64
}

func NewPolynomial(field *field.Field, coefficients []int64) *Polynomial {
	return &Polynomial{
		Field:        field,
		Coefficients: coefficients,
	}
}

func (poly *Polynomial) Degree() int {
	return len(poly.Coefficients) - 1
}

func (poly *Polynomial) EvaluateAt(point int64) int64 {
	powers := []int64{}
	for i := 0; i <= poly.Degree(); i++ {
		powers = append(powers, poly.Field.Exp(point, int64(i)))
	}
	return poly.EvaluatePowers(powers)
}

func (poly *Polynomial) EvaluatePowers(powers []int64) int64 {
	if len(powers) != len(poly.Coefficients) {
		// TODO: get rid of panic, implement proper error handling
		panic("powers should be the same size as co-efficients")
	}

	result := int64(0)
	for i := range poly.Coefficients {
		result = poly.Field.Add(
			result,
			poly.Field.Mul(poly.Coefficients[i], powers[i]),
		)
	}
	//println(result)
	return result
}

// TODO: Look into extracting the homorphic element into a separate package
// that implements the same arithmetic interface
func (poly *Polynomial) EvaluateEncryptedPowers(powers []int64) int64 {
	if len(powers) != len(poly.Coefficients) {
		// TODO: get rid of panic, implement proper error handling
		panic("powers should be the same size as co-efficients")
	}

	result := int64(1)
	for i := range poly.Coefficients {
		result = poly.Field.Mul(
			result,
			poly.Field.Exp(powers[i], poly.Coefficients[i]),
		)
	}
	return result
}
