package main

import (
	"fmt"
)

func  addre() func(int) int {
	sum := 0
	return func (v int)  int{
		sum += v
		return sum
	}
}

type iAdder func (int) (int, iAdder)

func addre2(base int) iAdder {
	return func (v int) (int, iAdder) {
		return base + v, addre2(base + v)
	}
}
func main()  {
	a := addre2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + ... + %d= %d\n", i,s)
	}
}