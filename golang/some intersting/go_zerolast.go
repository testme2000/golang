package main

import "fmt"

func main() {
	fmt.Println("Zero Last: Move all zero to last of array without changing the order")
	var arr []int = []int{2, 3, 0, 0, 4, 5, 0}
	var start int = 0
	end := 0
	var total_len int = len(arr)
	// Original array
	fmt.Println(arr)
	for start < total_len {
		if arr[start] == 0 {
			fmt.Println("Zero ")
			fmt.Println(start, end)
			fmt.Println(arr[start], arr[end])
			start += 1
		} else {
			fmt.Println("Non zero")
			fmt.Println(start, end)
			fmt.Println(arr[start], arr[end])
			temp := arr[start]
			arr[start] = arr[end]
			arr[end] = temp
			start += 1
			end += 1
		}
	}
	fmt.Println(arr)
}
