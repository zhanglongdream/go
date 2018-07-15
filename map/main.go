package main

import (
	"fmt"
)

func main()  {
	m := map[string] string{
		"name": "cc",
		"c": "bb",
	}
	m2 := make(map[string]int)
	var m3 map[string]int

	fmt.Println(m, m2, m3)

	for k, v := range m{
		fmt.Println(k, v)
	}
	mm1 := m["name"]
	mm2 , ok:= m["c"]
	fmt.Println(mm1, mm2, ok)
  mm21 , ok:= m["cc"]
fmt.Println(mm21,  ok)
}