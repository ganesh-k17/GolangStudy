/*

***Usin logrus framework for logging***
installing package: go get "github.com/sirupsen/logrus"


Supported Seviarity:

* Trace
* Debug
* Info
* Warn
* Error
* Fatal
* Panic
*/

package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
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

	log.Println("logging from logrus...")
	log.Panic("logging a panic") // To log panic
	log.Trace("logging a trace") // To log Trace

	log.SetLevel(log.DebugLevel)
	log.Debug("logging a debug ")
	/*
		To log debug.  Debug log wont print anything.  To get it printed we need to
		set the severity level to debug level as in previous line.
	*/
}
