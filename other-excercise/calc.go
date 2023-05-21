package main

import "fmt"

func calculate(a int, b int) []float64 {
	x := float64(a)
	y := float64(b)
	var results []float64 = []float64{x + y, x - y, x * y, x / y}
	// results = append(results, x+y)
	// results = append(results, x-y)
	// results = append(results, x*y)
	// results = append(results, x/y)
	return results
}

func main() {
	fmt.Println(calculate(20, 10))
	fmt.Println(calculate(700, 70))
}
