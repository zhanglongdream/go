package main

import "fmt"

func inset() func() int{
	i := 0
	return func() int{
		i += 1
		return i
	}
}
func main()  {
    a := inset()
    //更新 i的状态
    fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}
