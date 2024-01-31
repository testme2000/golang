package main

import "fmt"

func main() {
	fmt.Println("Fibo series with go : Recursive attempt")
	num := 6
	start := 0
	final := 0
	result := make([]int, num)
	fibo_recur(start, final, num, result)
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

func fibo_recur(start int, final int, num int, result []int) {

	if start == 0 {
		result[start] = final
		final = 1
	} else if start == 1 {
		result[start] = final
		final = 1
	} else {
		if start < num {
			prev := result[start-2]
			next := result[start-1]
			final = prev + next
			result[start] = final
		}
	}
	if start < num {
		start += 1
		fibo_recur(start, final, num, result)
	}
}
