package main

import (
	"os"
	"log"
	"fmt"
)

func main() {
	 file, err := os.Open("./lesson.go")

	 if err != nil {
		 log.Fatalln("EXIT", err)
	 }
	 defer file.Close()
	 data := make([]byte, 100)

	 count, err := file.Read(data)

	 if err != nil{
		 log.Fatalln("error")
	 }
	 fmt.Println(count, string(data))
}