package main

import (
    "fmt"
    "sort"
	// "github.com/markcheno/go-quote"
	// "github.com/markcheno/go-talib"
)

const (
    c1 = iota
    c2 
    c3
)

func main() {
    
    i := []int{1,2,3,4,5}
    s := []string{"d", "a", "f"}
    p := []struct{
        Name string
        Age int
    }{
        {"Nancy", 20},
        {"Vera", 40},
        {"Mike", 22},
    }

    fmt.Println(i, s, p)
    sort.Ints(i)
    fmt.Println(i)

    sort.Strings(s)
    fmt.Println(i, s)
}