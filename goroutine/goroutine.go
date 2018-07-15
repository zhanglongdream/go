package main

import (
	"time"
	"fmt"
)

func main()  {
	for i := 0; i < 10; i++ {
		go func (i int)  {
			for {
				fmt.Println("aa ", i)
				//交出控制权
				// runtime.Gosched()
			}	
		}(i)
	}
	time.Sleep(time.Minute)
}