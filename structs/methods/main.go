package main

import "fmt"


/*
  Go自动处理方法调用时的值和指针之间的转化。
  可以使用指针来调用方法来避免在方法调用时产生一个拷贝
  或者让方法能够改变接受的数据
*/
type rect struct {
	width, height int
}

func (r *rect) area() int {
	return  r.height * r.width
}

func (r *rect) perim() int  {
	return 2 * r.height + 2 * r.width
}
func main()  {
	r := rect{width: 100, height:4}

	fmt.Println(r)

	fmt.Println(r.area())

	fmt.Println(r.perim())
}
