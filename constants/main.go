package main

import "fmt"

const a = "const"
func main() {

	fmt.Println(a)

	const b = 520

	const d = b / 0.0001
	fmt.Println(int32(d))
}
