package main

import (
    "fmt"
    "encoding/json"
    "log"
    // "time"
)

type User struct {
    Id int
    Name string
    Email string
}

func main() {
    b := []byte(`{"id":1, "name":"tomoya", "email":"a@a.com"}`)
    p := new(User)

    if err := json.Unmarshal(b, &p); err != nil{
        log.Fatal(err)
    }
    fmt.Println(p.Id, p.Name, p.Email)

    v, _ := json.Marshal(p)
    fmt.Println(string(v))
}