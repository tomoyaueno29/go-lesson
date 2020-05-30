package main

import (
    "fmt"
    "encoding/json"
    "log"
    // "time"
)

type T struct{

}

type User struct {
    Id    int       `json:"id,string"`
    Name  string    `json:"name,omitempty"`
    Email string    `json:"email"`
    T     *T         `json:"T,omitempty"`
}

func (u User) MarshalJSON() ([]byte, error) {

    v, _ := json.Marshal(&struct{
        Name string
    }{
        
    })
}

func main() {
    b := []byte(`{"id":"1", "name":"", "email":"a@a.com"}`)
    p := new(User)

    if err := json.Unmarshal(b, &p); err != nil{
        log.Fatal(err)
    }
    fmt.Println(p.Id, p.Name, p.Email)

    v, _ := json.Marshal(p)
    fmt.Println(string(v))
}