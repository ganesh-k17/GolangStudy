package main

import "fmt"

type s1 struct {
	x int
}

type s2 struct {
	x int
}

func main() {
	c := s1{x: 5}
	c1 := s2{x: 5}
	if c == c1 { // it would cause error since different types,  cannot compare c == c1 (mismatched types s1 and s2)compilerMismatchedTypes
		fmt.Println("yes")
	}

	c2 := s1{x: 5}

	if c == c2 {
		fmt.Println("c and c2 are equal") // it works and give proper response as they are same type
	}
}
