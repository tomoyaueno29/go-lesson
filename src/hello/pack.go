package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main(){
	dir,_ := ioutil.TempDir("", "sa")
	fmt.Println(dir)
	defer os.RemoveAll(dir)
}