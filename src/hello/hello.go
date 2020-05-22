package main

import (
	"fmt"
)

func add(x int, y int) (int, int) {
	return x + y, x - y
}


func cal(price int, item int) (int) {

	result := price * item
	return result
}

func main() {
	result, result2 := add(20, 10)
	fmt.Println(result, result2)

	r3 := cal(100, 2)
	fmt.Println(r3)
}