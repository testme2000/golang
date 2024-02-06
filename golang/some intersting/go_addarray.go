package main

import (
	"fmt"
)

func main() {
	fmt.Println("Add two array, if sum is more than 9, use the first digit in location, carry the 10th ")
	fmt.Println("to next location")

	var list1 [5]int = [5]int{6, 8, 9, 1, 3}
	var list2 [5]int = [5]int{5, 2, 7, 8, 1}
	max := len(list1)
	if max < len(list2) {
		max = len(list2)
	}
	//var result = int(math.Max(float64(len(list1)),float64(100)))
	result := make([]int, max)
	var carry int = 0

	for count := range list1 {
		sum := list1[count]
		if count < len(list2) {
			sum += list2[count]
		}
		sum += carry
		oneplace := sum
		if sum > 9 {
			oneplace = sum % 10
			carry = sum / 10
		}
		result[count] = oneplace
	}
	fmt.Println(result)

}
