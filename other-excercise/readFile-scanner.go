package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filename string) ([]string, error) {
	// your code goes here
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func main() {
	lines, err := ReadLines("data.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, line := range lines {
			fmt.Println(line)
		}
	}

}
