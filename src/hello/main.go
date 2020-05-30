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
    Create time.Time
}

func main() {
    src := `
    {
    "Id":1,
    "Name":"tomoya",
    "Email":"aa@a.com",
    "time":"2020/4532928170223"
    }
    `

    u := new(User)

    err := json.Unmarshal([]byte(src), u)
    if err != nil{
        log.Fatal(err)
    }

    fmt.Println(u)
}