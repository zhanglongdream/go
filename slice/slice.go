package main

import (
	"fmt"
)
func updatesSlice(s []int)  {
	s[0] = 100
}
func main()  {
	arr := [...]int{0,1,2,3,4,5,6}

	fmt.Println("arr[2: 6]=", arr[2:6])
	fmt.Println("arr[2:]=", arr[2:])
  s1 := arr[:6]
	fmt.Println("s1=", s1)
	fmt.Println("arr[:]=", arr[:])
	updatesSlice(s1)
	fmt.Println(s1)

	fmt.Println(arr)
}