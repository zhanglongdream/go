package main

import (
	"fmt"
)

func printArray(arr *[5]int) {
	fmt.Println("arry 值传递")
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100
}

func main()  {
	var arr1 [5]int
	arr2 := [3]int{1,2,3}
	arr3 :=[...]int{2,3,4,5,67}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3, grid)

	for i := 0; i< len(arr3); i++ {
		fmt.Println(i)
	}
 fmt.Println("daaaa")
	for i ,  v:= range arr3 {
		fmt.Println(i, v)
	}
	for _,  v:= range arr3 {
		fmt.Println(v)
	}

	printArray(&arr1)
	fmt.Println(arr1)
	// printArray(arr2)
}