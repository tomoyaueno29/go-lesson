package main

import (
    "fmt"
	// "os"
	"io/ioutil"
	"log"
)

func main() {
	// 読み込み
	content, err := ioutil.ReadFile("pack.go")
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(string(content))
	// buf := make([]byte, 100)

	if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil{
		log.Fatalln(err)
	}


}