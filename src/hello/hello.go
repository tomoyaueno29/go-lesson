package main
 
import (
    "fmt"
)

func producer(ch chan int, i int){
    // Something
    ch <- i * 2
}

func main(){
    var wg sync.WaitGroup
    ch := make(chan int)

    for i := 0; i<10; i++{
        wg.Add(1)
        go producer(ch, i)
    }
}