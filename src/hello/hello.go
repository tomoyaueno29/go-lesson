package main
 
import (
    "fmt"
)
 
type Vertex struct{
    X, Y int
}

func (v Vertex) Plus() int{
    return v.X + v.Y
}


func (v *Vertex) Scale(i int){
    v.X = v.X * i
    v.Y = v.Y * i
}

// func Plus(v Vertex) int {
//     return v.X + v.Y
// }
 
func main(){
    v := Vertex{3, 4}
    // fmt.Println(Plus(v))
    fmt.Println(v.Plus())
}