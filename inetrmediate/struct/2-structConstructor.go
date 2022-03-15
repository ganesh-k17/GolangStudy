/*
https://gobyexample.com/structs
*/

package main

import "fmt"

type student struct {
	id   int
	name string
}

func newStudent(id int, name string) *student {
	return &student{id, name}
}

func main() {
	fmt.Println("hello")
	var myStudent = newStudent(1, "raja")
	fmt.Println(myStudent.id)
	fmt.Println(myStudent.name)
}
