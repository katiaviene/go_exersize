package main

import "fmt"

func main() {
	first := 0
	second := 1
	for i := 0; i <= 4; i++ {

		fmt.Println(first)
		first = first + second
		fmt.Println(second)
		second = first + second

	}
}
