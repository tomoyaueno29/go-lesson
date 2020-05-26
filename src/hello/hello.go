package main
 
import (
    "fmt"
)

type User struct {
    gender string
    age int

}

func (u User) String() string{
    return fmt.Sprintf("My name is %v\n", u.gender)
}

func main(){
   
    u := User{gender: "male", age: 22}
    fmt.Println(u.gender)
    fmt.Println(u.age)
    fmt.Println(u.String())
}