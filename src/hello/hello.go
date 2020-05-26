package main
 
import (
    "fmt"
    // "time"
)

func ping(pings chan<- string, msg string){

    pings <- msg
}

func pong(pings chan string, pongs chan string){

    msg := <-pings
    pongs <- msg
}


func main(){
    
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    var msg string = "Hello"
    
    go ping(pings, msg)
    go ping(pings, pongs)

    fmt.Println(<-pongs)
}