package main

import (
	"fmt"
	// "io/ioutil"
	// "net/http"
	"net/url"
	"log"
)

func main() {
	// resp, _ := http.Get("http://example.com")
	// defer resp.Body.Close()

	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	base, err := url.Parse("http://example.com")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(base)

	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println
}