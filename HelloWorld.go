package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	var num int64 = 1e9
	fmt.Println(num)
	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			fmt.Println(j, "is an even number")
		} else {
			fmt.Println(j, "is an odd number")
		}
	}
}
