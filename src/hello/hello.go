package main
 
import (
    "fmt"
)



func main(){
   
    s := Student{"Tomoya", 60.8, 90.0}
    fmt.Println(s.Name)
    fmt.Println(s.Math)
    fmt.Println(s.English)
    fmt.Println(s.Plus())
}