package main

import (
	"fmt"
	"strings"
)

func WordCount(s string, word string) int {
	// your code goes here
	return strings.Count(s, word)
}

func main() {
	count := WordCount("hello, Hello how have you been in helloworld", "hello")
	fmt.Println(count)
}
