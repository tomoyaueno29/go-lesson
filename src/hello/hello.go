package main
 
import (
    "fmt"
)
 
type Vertex struct{
    X, Y int
}


 
func main(){
    v := Vertex{3, 4}
    fmt.Println(v)
}