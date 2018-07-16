package main

import "fmt"

func main()  {
	a := [5]int{1,2,3,4,5}

	fmt.Println(a)
	for i, v := range a {
		fmt.Println(i, v)
	}

	//如果不想使用的话， 用_代替
}
