package main

import (
	"fmt"
	"time"
)

func work(c chan int) {
	for {
		n := <-c
		fmt.Println(n)
	}
}
func chanDemo() {
	c := make(chan int)
	go work(c)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}
func main() {
	chanDemo()
}
