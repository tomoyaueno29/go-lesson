package main
 
import (
    "fmt"
    "time"
)


func main(){
    
    tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisecond)

    for {
        isBreak := true
        select {

        case t := <- tick:
            fmt.Println("tick.", t)
        case <-boom:
            fmt.Println("BOOM!")
            // break
            // return 
        default:
            fmt.Println("    .")
            time.Sleep(50 * time.Millisecond)
        }
        if isBreak {
            break
        }
        
    }

    fmt.Println("############")
}