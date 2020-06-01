package main

import (
	"fmt"
	// "encoding/json"
	// "io/ioutil"
)

type Languages struct {
	Name, LangType string
}

func (l *Languages) Alter(name string, ltype string) {
	(*l).Name = name
	(*l).LangType = ltype
}

func New(name string, langtype string) *Languages {
	return &Languages{name, langtype}
}

func main() {
	 l := New("jap", "japanese")
	 l.Alter("en", "english")
	 fmt.Println(l.Name, l.LangType)
}