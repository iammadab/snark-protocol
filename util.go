package main

import "github.com/iammadab/snark-protocol/field"

func GenerateEncryptedPowers(point int64, degree int, generator int64, field *field.Field) []int64 {
	var encryptedPowers []int64
	for i := 0; i <= degree; i++ {
		power := IntPow(point, int64(i))
		encryptedValue := field.Exp(generator, power)
		encryptedPowers = append(encryptedPowers, encryptedValue)
	}
	return encryptedPowers
}
