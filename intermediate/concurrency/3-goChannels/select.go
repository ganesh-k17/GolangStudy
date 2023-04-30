package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go first(ch)
	go second(ch)

	select { // The cases are not deterministic select may pickup val1 or val2 randomly.
	case val1 := <-ch:
		fmt.Println(val1)
	case val2 := <-ch:
		fmt.Println(val2)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("end")
}

func first(ch chan int) {
	ch <- 1
}

func second(ch chan int) {
	ch <- 2
}
