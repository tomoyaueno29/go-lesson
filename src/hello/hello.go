package main
 
import (
    "fmt"
    "time"
)

func goroutine1(ch chan string){

    for {
        ch <- "packet from 1"
        time.Sleep(1 * time.Second)
    }
}

func goroutine2(ch chan string){

    for {
        ch <- "packet from 2"
        time.Sleep(1 * time.Second)
    }
}

func main(){
    
    c1 := make(chan string)
    c2 := make(chan string)

    go goroutine1(c1)
    go goroutine2(c2)
}