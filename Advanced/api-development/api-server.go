package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homepage)
	http.ListenAndServe("localhost:10000", nil)
	fmt.Println("end")
}

func homepage(writer http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(writer, "welcome to home")
	fmt.Println("Welcome endpoint hit...")
}
