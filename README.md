# Golang

This repo will have a collection of Golang concepts

## Go Build and Run

```go
go build hello.go
go run hello.go
```

Use go run when we want to treat a Go program like a script and run the source code immediately.

Most of the time we want to build a binary for later use. That’s where we use the go build command.

## Basics:

### Variable:

> var name type = expression

> var i int = 10

> var s string = "India"

> var i = 10

> var i int

> i := 10

### Scope of Variables:

- Scope of Golang Variables Defined by Brace Brackets.
- Goland uses lexical scoping based on code blocks to determine the scope of variables.
- Inner block can access its outer block defined variables.
- But outer block cannot access inner block defined variables.

```
package main

import (
	"fmt"
)

var s = "Japan"

func main() {
	fmt.Println(s)
	x := true

	if x {
		y := 1
		if x != false {
			fmt.Println(s)
			fmt.Println(x)
			fmt.Println(y)
		}
	}
	fmt.Println(x)
}
```

### Constants

```
const PRODUCT string = "Canada"
const PRICE = 500
```

### Multilple Constants as Block

```
const (
	PRODUCT  = "Mobile"
	QUANTITY = 50
	PRICE    = 50.50
	STOCK  = true
)
```

- Name of constants should follow the same rules as variable names, which means a valid constant name must starts with a letter or underscore, followed by any number of letters, numbers or underscores.
- By convention, constant names are usually written in uppercase letters. This is for their easy identification and differentiation from variables in the source code.

### Tips and tricks

http://www.golangbootcamp.com/book/tricks_and_tips

### Array and Slice difference:

| Array                    | Slice                         |
| ------------------------ | ----------------------------- |
| Arrays have fixed length | Slice don't have fixed length |
| myArray := [3]int{1,2,3} | mySlice := []int{1,2,3}       |
|                          |                               |

### Creating slice by make

```go
mySlice := make([]<datatype>, length, capacity)

mySlice := make([]int, 3, 4)
```
