package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func suns() int{
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return sum
}

func convertToBin(n int) string {
	result := ""
   for ; n > 0; n /= 2 {
		 lsb := n % 2
		 result = strconv.Itoa(lsb) + result
	 }
	 return result
}

func readFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func  forever()  {
	for {
		fmt.Println("abc")
	}
}
func main()  {
	fmt.Println(suns())

	fmt.Println(
		convertToBin(5),
		convertToBin(13),
	)

	readFile("abc.txt")
	forever()
}

//go  中没有 while