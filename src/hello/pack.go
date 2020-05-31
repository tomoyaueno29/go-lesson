package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	res, _ := http.Get("http://example.com")
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}