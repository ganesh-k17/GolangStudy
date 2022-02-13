/*
https://medium.com/kalamsilicon/hash-tables-implementation-in-go-48c165c54553
https://www.geeksforgeeks.org/golang-maps/

* Golang Maps is a collection of unordered pairs of key-value.
* It is widely used because it provides fast lookups and values that can retrieve, update or delete with the help of keys.
* It is a reference to a hash table.
* In the maps, a key must be unique, most of the built-in type can be used as a key like an int, float64, rune, string, comparable array and structure, pointer, etc
* In maps, the values are not unique like keys and can be of any type like int, float64, rune, string, pointer, reference type, map type, etc.
* The type of keys and type of values must be of the same type, different types of keys and values in the same maps are not allowed. But the type of key and the type values can differ.
* The map is also known as a hash map, hash table, unordered map, dictionary, or associative array.
In maps, you can only add value when the map is initialized if you try to add value in the uninitialized map, then the compiler will throw an error.
* As we know that maps are of reference type. So, when we assign an existing map to a new variable, both the maps still refer to the same underlying data structure. So, when we update one map it will reflect in another map.
*/

package main

import "fmt"

func main() {

	// An Empty map
	// map[Key_Type]Value_Type{}

	// Map with key-value pair
	// map[Key_Type]Value_Type{key1: value1, ..., keyN: valueN}

	myMap_1 := map[string]string{} // default initialisation for map

	if myMap_1 == nil {
		fmt.Println("True")
	} else {

		fmt.Println("False")
	}

	myMap_1["1"] = "ganesh" // adding values to map
	fmt.Println(myMap_1)

	myMap_2 := map[string]string{"1": "ganesh", "2": "raja"}

	fmt.Println(myMap_2)

	// using make

	myMap_3 := make(map[string]string)
	myMap_3["1"] = "vishnu"
	myMap_3["2"] = "Vinay"

	fmt.Println(myMap_3)

	// iterating over map

	for i, v := range myMap_3 {
		fmt.Printf("%s-%s", i, v)
		fmt.Println()
	}

	//If the key doesn’t exist in the given map, then it will return zero value of the map, i.e, nil. And if the key exists in the given map, then it will return the value related to that key.

	value_1 := myMap_3["90"]
	fmt.Println(value_1)

	pet_name, ok := myMap_3["90"] // retrieving the value based on key and validate it

	if ok {
		fmt.Println(pet_name)
	} else {
		fmt.Println("no key available")
	}

	delete(myMap_3, "1") // deleting a key from map

	fmt.Println(myMap_3)

	fmt.Println("end")
}
