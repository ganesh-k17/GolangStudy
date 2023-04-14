# Others

## Shorthand way

```go
package main
import ("fmt")

func main() {
    var s,t string = "foo", "bar"
    fmt.Println(s)
    fmt.Println(t)
}

/*
foo
bar
*/
```

```go
package main
import ("fmt")

func main() {
    var (
        s string = "foo"
        i int = 5
    )
    fmt.Println(s)
    fmt.Println(t)
}

/*
foo
bar
*/
```

```go
package main
import ("fmt")

func main() {
    name:== "Lisa"
    name = "Peter"
    fmt.Println(s)
    fmt.Println(t)
}

/*
Peter
*/
```

## Global variables

```go
package main
import ("fmt")

var name string = "Lisa"

func main(){
    fmt.Println(name)
}

/*
Lisa
*/
```

## Zero values

If variables are declared without values they are called Zero values. They would be assigned with default values based on the type it has been declared.

```string
Default values

bool -> false
int -> 0
float64 -> 0.0
string -> ""
pointers,function, interface, maps -> nil
```

## Single Input

```go
package main
import ("fmt")

func main(){
    var name string
    fmt.Print("Enter your name")
    fmt.Scanf("%s", &name)
    fmt.Println(name)
    fmt.Println("Hey, ",name)
}

/*
Enter your name: Priyanka
Hey Priyanka
*/
```

## Multiple Input

```go
package main
import ("fmt")

func main(){
    var name string
    var isMarried bool

    fmt.Print("Enter your name & Are you Married:")
    fmt.Scanf("%s %t", &name, &isMarried)
    fmt.Println(name)
    fmt.Println(name, isMarried)
}

/*
Enter your name & Are you Married: Priyanka true
Priyanka true
*/
```

## Scanf with return values

```go
package main
import ("fmt")

func main(){
    var a string
    var b int

    fmt.Print("Enter text and a number:")
    count, err := fmt.Scanf("%s %d", &a, &b)

    fmt.Println("count: ", count)
    fmt.Println("error: ",err)

    fmt.Println("a: ", a)
    fmt.Println("b: ", b)
}

/*
Enter text and a number: Priyanka yes
Count: 1
error: expected integer
a: priyanka
b: 0
*/
```

## Find type of variable

using %T

```go
package main

import "fmt"

func main() {
    var grades int = 42
	var message string = "hello world"
	var isCheck bool = true
	var amount float32 = 5466.54
	fmt.Printf("variable grades = %v is of type %T \n", grades, grades)
	fmt.Printf("variable message = '%v' is of type %T \n", message, message)
	fmt.Printf("variable isCheck = '%v' is of type %T \n", isCheck, isCheck)
	fmt.Printf("variable amount = %v is of type %T \n", amount, amount)
}

/*>>> go run main.go
variable grades = 42 is of type int
variable message ='hello world' is of type string
variable isCheck = 'true' is of type bool
variable amount = 5466.54 is of type float32
*/

```

Using reflect.TypeOf()

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("priyanka"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(46.0))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))
}

/*
>>> go run main.go
Type: int
Type: string
Type: float64
Type: bool
*/

```

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {

	var grades int = 42
	var message string = "hello world"
	fmt.Printf("variable grades=%v is of type %v \n", grades, reflect.TypeOf(grades))
	fmt.Printf("variable message='%v' is of type %v \n", message, reflect.TypeOf(message))
}

/*
>>> go run main.go
variable grades = 42 is of type int
variable message ='hello world' is of type string
*/

```

## Converting between types

`Type Casting`

• The process of converting one data type to another is
known as Type Casting.

• Data types can be converted to other data types, but this
does not guarantee that the value will remain intact.

Integer to float:

```go
package main

import "fmt"

func main() {

	var i int = 90
	var f float64 = float64(i)
	fmt.Printf("%.2f\n", f)
}

// >>> go run main.go
// 90.00
```

Float to integer

```go
package main

import "fmt"

func main() {

	var f float64 = 45.89
	var i int = int(f)
	fmt.Printf("%v\n", i)
}

// >>> go run main.go
// 45

```

strconv package:

`Itoa()`

• converts integer to string

• returns one value – string formed with the given integer.

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	var s string = strconv.Itoa(i) // convert int to string
	fmt.Printf("%q", s)
}

// >>> go run main.go
// “42”
```

Atoi()

• converts string to integer.

• returns two values – the corresponding integer, error (if any).

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "200"
	i, err := strconv.Atoi(s)
	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T", err, err)
}

// >>> go run main.go
// 200, int
// <nil>, <nil>
```

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {

	var s string = "200abc"
	i, err := strconv.Atoi(s)
	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T", err, err)
}

// >>> go run main.go
// 0, int
// strconv.Atoi: parsing "200a": invalid syntax, *strconv.NumError
```

Constants:

- Constants are untyped unless they are explicitly given a type at declaration.
- Allow for flexibility.

```go

// Untyped constants:

const age = 12
const h_name, h_age = "Hermione", 12

// typed constants:

const age int = 12
const h_name string = "Hermione"
```

Understanding Constants:

```go
package main

import "fmt"

func main() {
	const name = "Harry Potter"
	fmt.Printf("%v: %T \n", name, name)
}

// >>> go run main.go
// Error: cannot assign to name (declared const)
// name = "Hermione Granger

```

```go
package main
import "fmt"

func main() {
	const name
	fmt.Printf("%v: %T \n", name, name)
}


// >>> go run main.go
// missing value in const declaration,
// undefined: name
```
