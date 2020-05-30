package main

import (
	"fmt"
	"net/http"
	
)

func main() {
	resp, _ := http.Get("http://example.com")
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}