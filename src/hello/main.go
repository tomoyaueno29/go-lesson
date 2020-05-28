package main

import (
    "fmt"
)

// func reciever( ch <-chan int) {

//     for {
//         i := <-ch
//         fmt.Println(i)
//     }
// }

func main() {

    // ch := make(chan int)

    // go reciever(ch)

    // i := 0
    // for i < 100 {
    //     ch <- i
    //     i++
    // }

    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)
    fmt.Println(<- ch)
    fmt.Println(<- ch)
    fmt.Println(<- ch)
}