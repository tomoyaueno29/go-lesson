package main

import (
    "fmt"
    "time"
)

func main() {
    
    t := time.Now()
    fmt.Println(t.Format(time.RFC3339))
    fmt.Println(t.Year(), t.Day(), t.Hour(), t.Month())
}