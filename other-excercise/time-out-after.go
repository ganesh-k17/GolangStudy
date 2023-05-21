package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go setValue(ch)

	select {
	case value := <-ch:
		fmt.Println("value received is ", value)
	case <-time.After(1 * time.Second):
		fmt.Println("Time out after 1 second")
	}

	time.Sleep(2 * time.Second)
	fmt.Println("end")
}

func setValue(ch chan string) {
	// time.Sleep(3 * time.Second)
	ch <- "hello"
}
