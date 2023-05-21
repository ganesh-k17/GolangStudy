/*
We are building a Go application and we want to add logging to help you debug and troubleshoot issues. We have chosen to use the logrus package to accomplish this task.


Problem Statement:


The logrus package provides several functions for logging messages at different levels (e.g. Debug, Info, Warning, Error).
Our task is to use the logrus package to log messages at different levels. We should create a function called logMessage that takes in a level string and a message string as arguments.


If the level string is debug, the function should log the message using the Debug function.


If the level string is info, the function should log the message using the Info function.


If the level string is warning, the function should log the message using the Warning function.


If the level string is error, the function should log the message using the Error function.


If the level string is none of these, the function should log an error message using the Error function in the following format: -


"Invalid log level: " + level


A Go file is located at /root/code/logrus directory.
*/

package main

import log "github.com/sirupsen/logrus"

func logMessage(level string, message string) {
	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
		log.Debug("This is a debug message")
	case "info":
		log.Info("This is an info message")
	case "warning":
		log.Warning("This is a warning message")
	case "error":
		log.Error("This is an error message")
	default:
		log.Error("Invalid log level: " + level)
	}
}

func main() {
	log.SetLevel(log.DebugLevel)
	logMessage("debug", "This is a debug message")
	logMessage("info", "This is an info message")
	logMessage("warning", "This is a warning message")
	logMessage("error", "This is an error message")
	logMessage("invalid", "This is an invalid message")
}
