package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	/*
			$ go doc log setoutput
			package log // import "log"

			func SetOutput(w io.Writer)
		    SetOutput sets the output destination for the standard logger.
	*/
	log.SetOutput(file)

	log.Println("logging...")
}
