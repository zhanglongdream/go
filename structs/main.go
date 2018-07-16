package main

import "fmt"

/*
  go 的结构体是各个字段类型的集合

  使用这个语法创建了一个新的结构体元素
  可以在初始化一个结构体元素时指定字段名字
  省略的字段将被初始化为零值
  & 前缀生成一个结构体指针
*/
type person struct {
	name string
	age int
}
func main()  {
	fmt.Println(person{})

	fmt.Println(person{name: "zhanglong"})

	fmt.Println(person{"zhang", 1000})


	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "ll", age: 100}

	ss := &s
	fmt.Println(ss)

	ss.age = 1

	fmt.Println(ss, s)
}
