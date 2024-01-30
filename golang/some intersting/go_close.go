package main

import "fmt"

func main() {
	fmt.Println("Go: Closure feature")
	countodd := makeOddGenerator()
	for iter := 0; iter < 50; iter++ {
		fmt.Println(countodd())
	}
}

func makeOddGenerator() func() uint {
	count := uint(1)
	return func() (ret uint) {
		ret = count
		count += 2
		return
	}
}
