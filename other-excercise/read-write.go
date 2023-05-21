package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Open the input file
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()

	// Create a scanner to read the input file
	inputScanner := bufio.NewScanner(inputFile)

	// Initialize the sum to 0
	sum := 0

	// Iterate over the lines of the input file
	for inputScanner.Scan() {
		// Parse the current line as an integer
		n, err := strconv.Atoi(inputScanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		// Add the integer to the sum
		sum += n
	}

	// Check for errors while reading the input file
	if err := inputScanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	// Open the output file
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	// Write the sum to the output file as a string
	_, err = outputFile.WriteString(strconv.Itoa(sum))
	if err != nil {
		fmt.Println(err)
		return
	}
}
