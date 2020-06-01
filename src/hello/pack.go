package main

import (
	"fmt"
	"encoding/json"
	// "io/ioutil"
)

type Person struct {
	Name string	`json:"name"`
	Age int	`json:"age"`
	Nicknames []string	`json:"nicknames"`
}

func main() {
	b := []byte(`{"name":"tomoya", "age":22, "nicknames":["a", "b", "c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)
}