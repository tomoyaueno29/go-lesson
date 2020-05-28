package main

import (
    "fmt"
    "regexp"
	// "github.com/markcheno/go-quote"
	// "github.com/markcheno/go-talib"
)

func main() {
	// spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	// fmt.Print(spy.CSV())
	// rsi2 := talib.Rsi(spy.Close, 2)
    // fmt.Println(rsi2)
    
    match, _ := regexp.MatchString("a([a-z]+)e", "app0le")
    fmt.Println(match)
}