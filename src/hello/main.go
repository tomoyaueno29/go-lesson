package main

import (
    "fmt"
    "encoding/json"
)

type Person struct {
    Name      string    `json:"name"`
    Age       int       `json:"age,omitempty"`
    Nicknames []string  `json:"nicknames"`
}

func (p Person) MarshalJSON() ([]byte, error) {
    type Person2 struct{
        Name string
    }

    var p2 Person
    err := json
}

func main() {
    b := []byte(`{"name":"mike", "age":20, "nicknames":["a","b","c"]}`)
    var p Person
    if err := json.Unmarshal(b, &p); err != nil {
        fmt.Println(err)
    }
    fmt.Println(p.Name, p.Age, p.Nicknames)

    v, _ := json.Marshal(p)
    fmt.Println(string(v))
}