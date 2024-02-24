package main

import "fmt"

func main() {
	fmt.Println("Go: Turn taking array")

	limit := 10
	result := turn_double(limit)
	fmt.Println(result)
}

func turn_double(original int) []int {
	output := []int{}

	double_limit := original * 2
	first := 1
	for second := original; second < double_limit+1; second++ {
		output = append(output, first, second)
		first++
	}
	fmt.Println(output)

	return output
}
