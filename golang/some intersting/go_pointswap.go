package main

import "fmt"

func main() {
	x := 1
	y := 2

	fmt.Println("Swaping value using Pointer")
	fmt.Println("Original value", x, y)
	swap_it(&x, &y)
	fmt.Println("After swaping", x, y)
}

func swap_it(x *int, y *int) {
	temp := *y
	*y = *x
	*x = temp
}
