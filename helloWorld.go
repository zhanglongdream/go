package main

import (
	"fmt"
)

//包内部的变量
var a = 1

func variableZeroValue() {
	var a int 
	var s string 
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3,4
	var s string = "abc"
	fmt.Println(a, b, s)
}
func variableTypeDeduction() {
	var a, b, c, d = 1, true, 0, "a"
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	a, b, c, d := 1, true, 0, "a"
	fmt.Println(a, b, c, d)
}

func euler() {
	c := 3 + 4i
	fmt.Println(c)
}
func main() {
	fmt.Println("Hello world")
	fmt.Println(a)
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	euler()
}

// go 变量定义

// 使用var 关键字

// var a, b, c bool

// var s1, s2 string = "1", "11"

// 可放在函数内，或者直接放在包内部
// 使用var()集合定义变量

// 让编译器自动决定类型
// var a, b, c = 1, "2", true


// 使用 :=定义变量
// 只能在函数内部使用
//  a, b := 1, "11"