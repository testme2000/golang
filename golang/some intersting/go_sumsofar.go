package main

import "fmt"

func main() {
	fmt.Println("Go: Sum sofar")

	all_int := [5]int{1, 2, 3, 4, 5}

	sum_so_far(all_int[:])
	fmt.Println("Final Array {}", all_int)
}

func sum_so_far(array []int) {
	if len(array) < 2 {
		return
	}
	for idx, val := range array {
		if idx == 0 {
			continue
		}
		array[idx] = array[idx-1] + val
	}
}
