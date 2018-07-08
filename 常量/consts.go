package main


import (
	"math"
	"fmt"
)


func consts() {
	 const filename = "abc.txt"
	 const a,b = 3, 4
	 var c int
	 c = int(math.Sqrt(a * a + b * b))
	 fmt.Println(filename, c)
}

func enums()  {

	//使用变量来定义枚举类型
	const (
		cpp = 0
		java = 1
		python = 2
		golang = 3
	)
  const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func main()  {
	consts()
	enums()
}

//常量数值可作为各种类型使用

//