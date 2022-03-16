package main

import "fmt"

type student struct {
	id   int
	name string
}

func main() {
	fmt.Println("hello")
	var myStudent = student{1, "raja"}
	fmt.Println(myStudent.id)
	fmt.Println(myStudent.name)
}
