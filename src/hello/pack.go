package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	Name string
	Age int
	Nicknames []string
}

func main() {
	b := []byte(`{"name":"tomoya", "age":22, "nicknames":["a", "b", "c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	
}