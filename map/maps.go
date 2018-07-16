package main

import "fmt"

func main() {

	//创建了一个空的map
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)
	fmt.Println(len(m))
	delete(m, "k1")
	fmt.Println("map:", m)

	v, k := m["k2"]
	fmt.Println(v, k)

	a := map[string]string{"a": "string"}

	fmt.Println(a)

}
