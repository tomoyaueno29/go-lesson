package main

import (
    "fmt"
    // "io/ioutil"
)
func returnFunc() func(string) string{

    var store string

    return func(next string) string {

        s := store
        store = next
        return s
    }
}


func main() {

   x := returnFunc()
   fmt.Println(x("aaaa"))
   fmt.Println(x("bbbb"))
}