package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for i := 6; i < 9 ; i++ {
		fmt.Println(i)

	}
	for {
		fmt.Println("loop")
		break
	}
}
