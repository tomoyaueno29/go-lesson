package main

import (
	"fmt"
	// "io/ioutil"
	// "net/http"
	"net/url"
	"path"
)

func main() {
	// resp, _ := http.Get("http://example.com")
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	base, _ := url.Parse("http://foo.bar.jp/boo/?greeting=hello")
	copiedURL := *base
	fmt.Println(copiedURL.Path)
	copiedURL.Path = path.Join(copiedURL.Path, "./bee")
	
	fmt.Printf("before: %s\n", base)
    fmt.Printf("after : %s\n", &copiedURL)
}