package main

import "fmt"

func main() {
	fmt.Println("Tower of Hanoi: Using golang ")
	total := 3
	tower(total, "A", "C", "B")
}

func tower(count int, source, destination, intermediate string) {
	if count == 1 {
		fmt.Println("Move 1 disc from ", source, " to ", destination)
		return
	}
	tower(count-1, source, intermediate, destination)
	fmt.Println("Move disc ", count, " from ", source, " to ", destination)
	tower(count-1, intermediate, destination, source)
}
