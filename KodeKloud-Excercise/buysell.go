package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)
	go sell(ch)
	time.Sleep(3 * time.Second)
	go buy(ch)

	time.Sleep(1 * time.Second)

	fmt.Println("end")
}

func sell(ch chan string) {
	ch <- "furniture"
	fmt.Println("selling furniture")
}

func buy(ch chan string) {
	fmt.Println("Waiting for buy")
	value := <-ch
	fmt.Println("Received...", value)
}
