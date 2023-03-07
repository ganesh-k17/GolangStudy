# Types

- String
- Number => Integer, Float
- Boolean (true, false)
- Array & Slices
  Arrays => List or sequence of elements with same types
  Slices => Flexible array with more control on memory allocation.
- Map => Collection of key, value pairs.

## Static vs Dynamic typed languages

Static typed:

- Compiler throws an error when types are used incorrectly.
- Examples: c++, Java, C#
- Advantages
  - Better performance.
  - Bugs can often be caught by a compiler
  - Better data integrity

Dynamic typed:

- Compiler does not enforce the type system.
- Examples: Python, Javascript.
- Advantages:
  - Faster to write code
  - Generally, less rigid.
  - Shorter learning curve

## In Golang?

- Go has a concept of types that is either explicitly declared by a programmer or is inferred by the compiler.
- It's a fast, `statically typed`, compiled language that `feels like a dynamically typed`, interpreted language.

## Integers

- uint8 (8 bits or 1byte)
- uint16 (16 bits or 2 bytes)
- uint32 (32 bits or 4 bytes)
- uint64 (64 bits or 8 bytes)
- int8 (8 bits or 1 byte)
- int16 (16 bits or 2 bytes)
- int32 (32 bits or 4 bytes)
- int64 (64 bits or 8 bytes)
- int (4 bytes for 32 bit machines, 8 bytes for 64 bit machines)

## Boolean (1 byte)

- bool (true, false)

## Arrays, Slices and Maps

- Array & Slices

  - [1, 2, 4, 9]
  - ["",""]
  - [7.0, 9.43, 0.65]

- Maps

  - "x" -> 30
  - 1 -> 100
  - "key" -> "value"

## Variables

Syntax:

```go
var <variable name> <data type> = <value>

var s string = "hello world"
var i int = 100
var b bool = false
var f float64 = 77.90

```

## Printing Variables

```go
fmt.Print("Hello World")

var city string = "Kolkata"
fmt.Print(city)


var name string = "KodeKloud"
var user string = "Harry"
fmt.Print("Welcome to ", name, ", ", user)   // Welcome to KodeKloud, Harry

var name string = "KodeKloud"
var user string = "Harry"
fmt.Print(name, "\n")
fmt.Print(user)

/*
KodeKloud
Harry
*/

fmt.Println(name)
fmt.Println(user)

/*
KodeKloud
Harry
*/

```

## String formatting

```go
fmt.Printf("Template string %s", object args(s))
```

- Print - format specifier

% v -> formats the value in default format.

```go
var name string = "KodeKloud"
fmt.Printf("Nice to see you here, at %v", name)

// Nice to see you here, at KodeKloud
```

%d => formats decimal integers

```go
var grades int = 42
fmt.Printf("Marks: %d", grades)

// Marks: 42
```

Example:

```go
package main
import ("fmt")

func main() {
  var name string = "Joe"
  var i int = 78
  fmt.Printf("Hey, %v! You have scored %d/100 in Physics", name, i)
}

// Hey, Joe! You have scored 78/100 in Physics
```

| Verb | Description                           |
| ---- | ------------------------------------- |
| %v   | default format                        |
| %T   | type of the value                     |
| %d   | integers                              |
| %c   | character                             |
| %q   | quoted characters/string              |
| %s   | plain string                          |
| %t   | true or false                         |
| %f   | floating numbers                      |
| %2f  | floating numbers upto 2 decimal parts |
