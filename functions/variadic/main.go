package main

import "fmt"

//可变参数函数， 可以用任意数量的参数调用

func sum (sums ...int) {
	fmt.Println(sums)

	sum := 0
	for _, v := range sums {
		sum += v
	}

	fmt.Println(sum)
}

func addString(a ...string)  {
	fmt.Println(a)
	s := ""
	for _, v := range a {
		s += v
	}
	fmt.Println(s)
}
func main()  {
	sum(1,2,3)
	sum(1,2,3,4,5)
	addString("1", "2", "dayu")
}
