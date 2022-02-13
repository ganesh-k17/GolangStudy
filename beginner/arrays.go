/*

* An array is a fixed-length sequence that is used to store homogeneous elements in the memory
* Due to their fixed length array are not much popular like Slice in Go language
* The elements of the array are indexed by using the [] index operator with their zero-based position, means the index of the first element is array[0] and the index of the last element is array[len(array)-1]
* In Go language, the array type is one-dimensional.
* The length of the array is fixed and unchangeable.
* You are allowed to store duplicate elements in an array
* Using shorthand declaration: In Go language, arrays can also declare using shorthand declaration. It is more flexible than the above declaration.
  Syntax:

	array_name:= [length]Type{item1, item2, item3,...itemN}

* As we already know that arrays are 1-D but you are allowed to create a multi-dimensional array. Multi-Dimensional arrays are the arrays of arrays of the same type

  Array_name[Length1][Length2]..[LengthN]Type

  array_name:= [2][2]string{ {"",""}, {"",""} }

* In an array, if an array does not initialized explicitly, then the default value of this array is 0

  var myarr[2]int
  fmt.Println("Elements of the Array :", myarr)  // output will be[0,0]

* In an array, if ellipsis ‘‘…’’ become visible at the place of length, then the length of the array is determined by the initialized elements.

   myarray:= [...]string{"GFG", "gfg", "geeks", "GeeksforGeeks", "GEEK"}

* Accessing integer array,

  for x:=0; x < len(myarray); x++{
	fmt.Printf("%d\n", myarray[x])
  }

* In Go language, an array is of value type not of reference type. So when the array is assigned to a new variable, then the changes made in the new variable do not affect the original array
* In an array, if the element type of the array is comparable, then the array type is also comparable. So we can directly compare two arrays using == operator.

arr1:= [3]int{9,7,6}
arr2:= [...]int{9,7,6}
arr3:= [3]int{9,5,3}

here arr1 and arr2 are equal

fmt.Println(arr1==arr2)  // true
fmt.Println(arr2==arr3)  // false
fmt.Println(arr1==arr3)  // true

*/

package main

import "fmt"

func main() {

	// var employeeNames = [3]string{"raja", "kumar", "vijay"}

	// employeeNames := [3]string{"raja", "kumar", "vijay"}

	// Creating array by deciding the size based on the number of elements

	employeeNames := [...]string{"raja", "kumar", "vijay"} // ... is the spread operarator

	for i, v := range employeeNames {
		fmt.Printf("%d", i)
		fmt.Println(v)
	}

	fmt.Print("end")
}
