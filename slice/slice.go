package main

import (
	"fmt"
)
func updatesSlice(s []int)  {
	s[0] = 100
}
func main()  {
	arr := [...]int{0,1,2,3,4,5,6}
  
	s11 := arr[2:6]
	s22 := s11[3:5]
	fmt.Printf("s11 = %v , len(s11) = %d, cap(s11)=%d\n",
						 s11, len(s11), cap(s11))
	fmt.Printf("s22 = %v , len(s22) = %d, cap(s22)=%d\n",
						 s22, len(s22), cap(s22))
  fmt.Println("============")
	fmt.Println("arr[2: 6]=", arr[2:6])
	fmt.Println("arr[2:]=", arr[2:])
  s1 := arr[:6]
	fmt.Println("s1=", s1)
	fmt.Println("arr[:]=", arr[:])
	updatesSlice(s1)
	fmt.Println(s1)

	fmt.Println(arr)


}

// slice可以向后扩展，不可以向前扩展
//s[i]不可以超越len(s)，向后扩展不可以超过底层数组cap(s)