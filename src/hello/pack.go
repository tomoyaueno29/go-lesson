package main

import (
	"fmt"
	// "io/ioutil"
	"net/http"
	"net/url"
	// "path"
)

func main() {
	// resp, _ := http.Get("http://example.com")
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("If-None-Match", "W/wywyeyerf23")
	q := req.URL.Query()
	fmt.Println(q)
	q.Add("c", "3&%")
	fmt.Println(q.Encode())

	req.URL.RawQuery = q.Encode()
	// a=1&b=2&c=3%26%25


}