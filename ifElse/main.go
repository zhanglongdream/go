package main

import "fmt"

func main() {
	if 7 % 2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
	var a int = 10
	if  a > 10 {
		fmt.Println("max")
	} else if   0 < a && a < 9 {
		fmt.Println("deng")
	} else {
		fmt.Println("min")
	}
}
