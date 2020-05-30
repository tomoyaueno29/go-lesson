package main

import (
    "fmt"
    "encoding/json"
)

type User struct {
    Id int
    Name string
    Email string
}

func main() {
    u := new(User)
    u.Id = 1
    u.Name = "tomoya"
    u.Email = "xxx@a.com"

    bs, _ := json.Marshal(u)
    fmt.Println(string(bs))
}