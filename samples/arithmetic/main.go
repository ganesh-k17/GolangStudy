package main

import (
	"fmt"

	"github.com/arithmetic/addition"
	"github.com/pborman/uuid"
)

func main() {
	uuid := uuid.NewRandom()
	fmt.Println("addition - ", addition.Add(5, 6))
	fmt.Println("uuid - ", uuid)
	fmt.Println("end")
}
