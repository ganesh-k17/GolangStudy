/*

* In Go language slice is more powerful, flexible, convenient than an array
* It is a lightweight data structure.
* Slice is a variable-length sequence which stores elements of a similar type, you are not allowed to store different type of elements in the same slice.
* It is just like an array having an index value and length, but the size of the slice is resized they are not in fixed-size just like an array. Internally, slice and an array are connected with each other, a slice is a reference to an underlying array
* It is allowed to store duplicate elements in the slice. The first index position in a slice is always 0 and the last one will be (length of slice – 1).
* A slice is declared just like an array, but it doesn’t contain the size of the slice. So it can grow or shrink according to the requirement.

 */

package main

import "fmt"

func main() {
	var mySlice = []string{}
	fmt.Println(mySlice)                         // []
	mySlice = append(mySlice, "raja")            // appending an item
	fmt.Println(mySlice)                         // [raja]
	mySlice = append(mySlice, "vijay", "vikram") // appending multiple items in slice
	fmt.Println(mySlice)                         // [raja vijay vikram]
	// mySlice = append(mySlice[:0], "vikram")

	fmt.Println(mySlice[:1], mySlice[1:]) // printing all the data before 1 index (:1) and all the data from 1 index (1:)
	fmt.Println(mySlice[:2], mySlice[2:]) // printing all the data before 2 index (:2) and all the data from 2 index (2:)

	mySlice = append(mySlice[:2], mySlice[2:]...) // append (combine) all the items which are before 2nd index and all the items which are from 2nd index
	fmt.Println(mySlice)

	mySlice = append(mySlice[:1], mySlice[2:]...) // removing the item vijay in [raja vijay vikram]
	fmt.Println(mySlice)

	mySlice = append(mySlice, "Vishnu") // adding an item to the existing slice
	fmt.Println(mySlice)
	fmt.Println("end")
}
