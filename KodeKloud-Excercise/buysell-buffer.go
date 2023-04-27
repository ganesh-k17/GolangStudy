package main

import (
	"fmt"
	"sync"
	// "time"
)

func main() {
	ch := make(chan string, 5)
	var wg sync.WaitGroup
	println("capacity of the channel", cap(ch)) // to calcualte capacity of the channel

	wg.Add(2)
	sell(ch, &wg)
	buy(ch, &wg)

	wg.Wait()

	// time.Sleep(1 * time.Second)
	fmt.Println("end")
}

func buy(ch chan string, wg *sync.WaitGroup) {
	// value := <-ch
	fmt.Println("received ", <-ch)
	fmt.Println("received ", <-ch)
	fmt.Println("received ", <-ch)
	wg.Done()
}

func sell(ch chan string, wg *sync.WaitGroup) {
	ch <- "furniture"
	ch <- "Chairs"
	ch <- "Tables"
	fmt.Println("send furniture, chair, tables")
	wg.Done()
}
