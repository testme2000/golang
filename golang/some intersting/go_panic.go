package main

import "fmt"

func main() {

	defer func() {
		str := recover()
		fmt.Println(str)
		tryit()
	}()
	fmt.Println("Everything is smooth now")
	panic("Meku created some issue")
}

func tryit() {
	fmt.Println("Within tryit function")
}
