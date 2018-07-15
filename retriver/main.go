package main

import (
	"fmt"
	"./mock"
	"./real"
)

type retriver interface {
	Get(url string) string
}
func download(r retriver) string {
	return r.Get("http://www.baidu.com")
}
func main()  {
	var r retriver
	r = mock.Retriver{"this is mock"}
	fmt.Printf("%T %v\n", r, r)
	r = real.Retriver{}
	fmt.Printf("%T %v\n", r, r)
	// fmt.Println(download(r))
}