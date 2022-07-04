package main

import "fmt"

// Consider the following two functions which sum values in a map
// These functions use different map types, but their underlying
// code is the same.

// SumInts adds together the values in a map of string to int64.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values in a map of string to float64
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	// With generics one function can generate these sums instead of two.

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	// This function signature has two type parameters (inside the square brackets), K and V
	// It has one argument (inside the round brackets), m which is of type map[K]V/
	// it also returns a value of type V

	// The comparable constraint is pre-declared in Go it allows any type whose values
	// may be used as an operand of the comparision operators == and !=
	// Go requires that map keys be comparable. So declaring K as comparable is necessary
	// so that K can be used as the key in the map variable. It also ensures that calling
	// code uses an allowable type for map keys.

	// The V type parameter is a constraint that is a union of two types int64 and float64
	// using | specifies a union of the two types meaning this constraint allows either type.
	// Either int64 or float64 will be permitted by the compiler as an argument for calling
	// this function.

	// It is known that map[K]V is a valid map type because K is a comparable type. If K had
	// not be declared comparable, the compiler would reject the reference to map[K]V

	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Type constraints can also be declared
// This is an example of declaring a number interface which is a union of int64 and float64
type Number interface {
	int64 | float64
}

// After declaring a type constraint it can be used as a type parameter when declaring
// a generic function, for example:

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
