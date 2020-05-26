package main
 
import (
    "fmt"
    "time"
)

type Counter struct {

    v map[string]int
    mux sync.Mutex

}

func (c * Counter) Inc(key string){
    c.mux.Lock()
    c.v[key]++
}

func main(){
    
    c := make(map[string]int)
    go func(){  
        for i := 0; i<10; i++{

            c["key"] += 1
        }
    }()

    go func(){  
        for i := 0; i<10; i++{

            c["key"] += 1
        }
    }()
    time.Sleep(1 * time.Second)
    fmt.Println(c, c["key"])
}