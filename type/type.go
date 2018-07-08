package main

import (
	"reflect"
	"fmt"
)

func  main()  {
	var a = 1
	b := string(a)
	fmt.Println(reflect.TypeOf(a) )
	fmt.Println(reflect.TypeOf(b) )
}

// 在go中类型转换都是强制的