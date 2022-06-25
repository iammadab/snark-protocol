package main

import "github.com/iammadab/snark-protocol/field"

// TODO: Seperate into different packages
func main() {
	a := field.NewField(7)
	println(a.Exp(2, 3))
}
