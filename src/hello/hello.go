package main

import (
	"fmt"

)

func main() {
	
	for i := 0; i<10; i++{
		if i==3{
			continue
		}

		// fmt.Println(i)
	}

	var sum int = 1
	for sum<10 {
		sum+=sum
		fmt.Println(sum)
	}
	fmt.Println(sum)

	// 無限ループ
	for {
		fmt.Println("hello")
	}
}