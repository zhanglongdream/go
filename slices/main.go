package main

import (
	"fmt"
	"reflect"
)
func aa(a []int) {
	a[0] = 1
}
func main() {
	//创建一个空的map
	s := make([]string, 3)

	fmt.Print("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)

	fmt.Println(s[1], len(s))

	//copy

	c := make([] string, len(s))
	copy(c, s)

	fmt.Println(c)

	a := c[1:]
	fmt.Println(a)

	t := []string{"a", "b", "c"}

	fmt.Println(t, reflect.TypeOf(t))

	q := make([][]int, 3)

	fmt.Println(q)
    bb := []int {111,2,3}
    fmt.Println(bb)
	aa(bb)
    fmt.Println(bb)
}
