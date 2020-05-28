package main

import (
    "fmt"
    "time"
    // "sort"
	// "github.com/markcheno/go-quote"
	// "github.com/markcheno/go-talib"
)

func longProcess(ch chan string){

    fmt.Println("run")
    time.Sleep(2 * time.Second)
    fmt.Println("finish")
    ch <- "result"
}

func main() {
    ch := make(chan string)
    ctx := context.Background()

    go longProcess(ch)

    for {
        select {
        
        case <- ch:
            fmt.Println("seccess")
            return
        }
    }
}