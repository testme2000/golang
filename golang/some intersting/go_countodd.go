package main

import "fmt"

func main() {
	fmt.Println("Count Odd between two numbers")
	fmt.Println("Count Odd between both odd")
	result := countOdd(3, 7)
	fmt.Println(result)
	fmt.Println("Count Odd between both even/odd and even/even")
	result = countOdd(2, 8)
	fmt.Println(result)
	result = countOdd(2, 7)
	fmt.Println(result)
}

func countOdd(low, high int) int {
	var count int = 0

	if low%2 == 1 && high%2 == 1 {
		count = ((high - low) / 2) + 1
	} else {
		count = (high - low) / 2
	}
	return count
}
