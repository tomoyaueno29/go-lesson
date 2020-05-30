package main

import (
    "fmt"
    "encoding/json"
    "log"
    "time"
)

type User struct {
    Id int
    Name string
    Email string
}

func main() {
    b := []byte(`{"Id":1, "Name":"tomoya", "Email":"a@a.com"}`)

    
}