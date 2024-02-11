package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Go: With Palindrome sceanrio")
	result := is_palindrom("apple")
	fmt.Println("Result is palindrome ", result)
}

func is_palindrom(input string) bool {
	var result bool = true
	var finalstring = strings.ToLower(input)
	var end int = len(input) - 1
	for _, val := range finalstring {
		var forcomp byte = byte(val)
		if forcomp != input[end] {
			result = false
			break
		}
		end--
	}

	return result
}
