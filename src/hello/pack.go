package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
)

func main() {
	// resp, _ := http.Get("http://example.com")
	// defer resp.Body.Close()

	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(body)
	// fmt.Println(string(body))

	base, _ := url.Parse("http://example.com/")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("acabckava", "Wcdcnalc")
	q := req.URL.Query()
	q.Add("c", "3")
	req.URL.RawQuery = q.Encode()
	fmt.Println(q.Encode())

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}