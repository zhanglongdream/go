package main

import (
	"fmt"
)

func printSlice(s []int){
	fmt.Printf("%v, len = %d, cap=%d\n",s, len(s), cap(s))
}
//添加元素时如果超越cap,系统会重新分配更大的底层数组
//由于值传递的关系，必须接受append的返回值
func main()  {
	s1 := [...]int {0,1,2,3,4,56}
	s2 := s1[:]
	s3 := append(s2, 111)
	fmt.Println(s3)

	var a []int //Zero value for slice is nil
	for i := 0; i< 100;i++ {
		// printSlice(a)
		a = append(a, 2 * i + 1)
	}
	// fmt.Println(a)
	ab := []int{11,2,3,4}
	printSlice(ab)
	b := make([]int, 16)
	c := make([]int, 10, 32)

	printSlice(b)
	printSlice(c)

	copy(c, ab)
	printSlice(c)

	c = append(c[:3], c[4:]...)
	fmt.Println(c)
}