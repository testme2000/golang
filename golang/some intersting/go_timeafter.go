package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			// c2 <- "from 2"
			// time.Sleep(time.Second * 2)
			<-time.After(time.Second * 3)
			c2 <- "From time after"
			//fmt.Println(aftertime)
		}
	}()

	go func() {
		for {
			select {
			case <-time.After(time.Second * 2):
				fmt.Println("from 1")
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)

}
