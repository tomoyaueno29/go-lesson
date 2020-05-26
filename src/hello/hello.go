package main
 
import (
    "fmt"
    // "time"
)

func ping(pings chan<- string, msg string){

    pings <- msg
}

func 


func main(){
    
    pings := make(chan string, 1)
    pongs := make(chan string, 1)

    go ping(pings, msg)
    go ping(pongs, msg)

    message <- "hello"
    message <- "world"

    fmt.Println(<-message)
    fmt.Println(<-message)
}