package main

import (
	"github.com/iammadab/snark-protocol/field"
	"math"
)

func GenerateEncryptedPowers(point int64, degree int, generator int64, field *field.Field) []int64 {
	var encryptedPowers []int64
	for i := 0; i <= degree; i++ {
		power := IntPow(point, int64(i))
		encryptedValue := EncryptValue(power, generator, field)
		encryptedPowers = append(encryptedPowers, encryptedValue)
	}
	return encryptedPowers
}

func ShiftEncryptedPowers(encryptedPowers []int64, shift int64, field *field.Field) []int64 {
	var shiftedPowers []int64
	for i := 0; i < len(encryptedPowers); i++ {
		shiftedPower := field.Exp(encryptedPowers[i], shift)
		shiftedPowers = append(shiftedPowers, shiftedPower)
	}
	return shiftedPowers
}

func EncryptValue(point int64, generator int64, field *field.Field) int64 {
	return field.Exp(generator, point)
}

// IntPow performs integer exponentiation
func IntPow(a, b int64) int64 {
	return int64(math.Pow(float64(a), float64(b)))
}
