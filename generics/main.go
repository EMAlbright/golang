package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumInts(m map[string]int64) int64 {
	var s int64

	// loop over map m and sum the values
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64

	for _, v := range m {
		s += v
	}
	return s
}

/*
Go requires that map keys be comparable. So declaring K as comparable is necessary so
you can use K as the key in the map variable.
It also ensures that calling code uses an allowable type for map keys
*/
// If we hadn’t declared K comparable, the compiler would reject the reference to map[K]V

func SumIntsAndFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	/*
		var mInt map[string]int64
		mInt = make(map[string]int64) // Now it's an empty map, not nil

		mInt["First"] = 5
		mInt["Second"] = 10
	*/

	mInt := map[string]int64{
		"First":  5,
		"Second": 5,
	}

	mFloat := map[string]float64{
		"First":  50.5,
		"Second": 75.2,
	}

	fmt.Printf("Non-Generics Sums: %v and %v\n",
		SumInts(mInt),
		SumFloats(mFloat))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsAndFloats(mInt),
		SumIntsAndFloats(mFloat))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumNumbers(mInt),
		SumNumbers(mFloat))

}
