package main

import "fmt"

func zaroval(n int) {
	n  = 0
}

func zeroptr(n *int) {
	*n  = 0
}
func main()  {
    a := 1
	zaroval(a)
   fmt.Println(a)

    zeroptr(&a)
    fmt.Println(a)

    //在 zaroval 传入的值是 a的拷贝
    // zeroptr 传入的是一个指针指向了 a内存中的值
}
