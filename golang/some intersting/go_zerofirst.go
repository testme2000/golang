package main

import "fmt"

func main() {
	fmt.Println("Zero First: Move all zero to start of array without changing the order")
	var arr []int = []int{2, 3, 0, 0, 4, 5, 0}
	var start int = len(arr) - 1
	end := len(arr) - 1
	// Original array
	fmt.Println(arr)
	for start >= 0 {
		if arr[start] == 0 {
			start -= 1
			fmt.Println("Start", start)
		} else {
			fmt.Println("Start", start)
			fmt.Println("End", end)
			fmt.Println(arr)
			temp := arr[start]
			arr[start] = arr[end]
			arr[end] = temp
			start -= 1
			end -= 1
			fmt.Println(arr)
		}
	}
	fmt.Println(arr)
}
