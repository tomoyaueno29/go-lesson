package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {

    content, err := ioutil.ReadFile("main.go")
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println(string(content))

    if err := ioutil.WriteFile("ioutil temp.go", content, 0666); err != nil{
        log.Fatalln(err)
    }
}