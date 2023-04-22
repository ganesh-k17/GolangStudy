/*
* user-defined data type
* a structure that groups together data elements
* provide a way to reference a series of grouped values through a single variable name.
 */

package main

import "fmt"

type student struct {
	id   int
	name string
}

func main() {
	fmt.Println("hello")
	var myStudent = student{id: 1, name: "raja"} // student{1, "raja"}
	fmt.Println(myStudent.id)
	fmt.Println(myStudent.name)
}
