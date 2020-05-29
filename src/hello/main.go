package main

import (
    "fmt"
    // "net/http"
    "net/url"
    // "io/ioutil"
)

func main() {
    // resp, _ := http.Get("http://ezample.com")
    // defer resp.Body.Close()
    // body, _ := ioutil.ReadAll(resp.Body)
    // fmt.Println(string(body))

    base, _ := url.Parse("http://example.com")
    reference, _ := url.Parse("/test?a=1&b-=2")
    endpoint := base.ResolveReference(reference).String()
    fmt.Println(endpoint)
}