package main

import (
	"fmt"
)

type Human interface {
    Say() string
}

type Person struct {
    Name string
}   

type Dog struct {
    Name string
}

func (p *Person) Say() string {
    (*p).Name = "Mr." + p.Name
    fmt.Println(p.Name)
    return p.Name
}

func Drivecar(human Human){
    if human.Say() == "Mr.Mike"{
        fmt.Println("Run")
    }else{
        fmt.Println("Get out")
    }
}

func main() {
    
    var mike Human = &Person{"Mike"}
    var dog Dog = Dog{"dog"}
    Drivecar(mike)
    Drivecar(dog)
}