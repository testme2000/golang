package main

import "fmt"

func main() {
	fmt.Println("Fibo series with go : Recursive attempt")
	result := fibo(6)
	fmt.Println(result)
}

func fibo(num int) []int {
	result := make([]int, num)
	if num == 0 {
		result = append(result, 0)
	} else if num == 1 {
		result = append(result, 1)
	} else {
		result[1] = 1
		fmt.Println(result)
		prev := 1
		next := 1
		final := 1
		for count := 2; count < num; {
			result[count] = final
			fmt.Println(result)
			final = prev + next
			prev = next
			next = final
			count++
		}
	}

	return result
}
