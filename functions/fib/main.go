package main

import (
	"fmt"
)

func fbnq() func() int {
	a, b := 0, 1
	return func ()  int{
			a ,b = b , a + b
			return b
	}
}
func main()  {
	s := fbnq()
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())

	fmt.Println(s())
	fmt.Println(s())
	fmt.Println(s())
}