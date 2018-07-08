package main

import (
	"fmt"
	"io/ioutil"
)
// if的条件里可以赋值
//if的条件里赋值的变量作用域就在if语句里面

func ifFun() {
  const filename = "abc.txt"
	// contents, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }
  if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func switchFunc(score int) string {
	g := ""
  switch  {
	case score < 60:
		g = "F"
	case score < 80:
		g = "c"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	default:
		fmt.Printf("%d\n ", score)
	}
	return g
}
func main()  {
	ifFun()
	fmt.Println(switchFunc(80))

}