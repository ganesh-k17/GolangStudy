/*
Our task is to create a new error type called APIError that implements the error interface.
The APIError type should have the following fields:


Message: A string that describes the error.


Code: An integer that represents the error code returned by the API.


Details: A map[string]interface{} that contains additional error details.


We should implement the Error method for the APIError type, which should return the error message. Also, we should implement a function called NewAPIError that creates a new instance of the APIError type.


A Go file is located at /root/code/infra directory.
*/

package main

import (
	"fmt"
)

type APIError struct {
	Message string
	Code    int
	Details map[string]interface{}
}

func (e *APIError) Error() string {
	return e.Message
}

func NewAPIError(code int, message string, details map[string]interface{}) error {
	return &APIError{
		Message: message,
		Code:    code,
		Details: details,
	}
}

func main() {
	err := NewAPIError(400, "Bad request", map[string]interface{}{
		"field": "username",
		"error": "cannot be empty",
	})
	fmt.Println(err)
}
