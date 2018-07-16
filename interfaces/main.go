package main

import (
	"fmt"
	"github.com/pkg/errors"
)

//接口是方法特征的命名集合
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width + r.height
}

func (r rect) perim() float64 {
	return  2 * r.width + 2 *r.height
}

func (r circle) area() float64 {
	return r.width + r.height
}

func (r circle) perim() float64 {
	return  2 * r.width + 2 *r.height
}

func measure(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main()  {

    r := rect{width: 10, height: 100}
    s := circle{width: 20 , height: 100}
	measure(r)
	measure(s)
}
