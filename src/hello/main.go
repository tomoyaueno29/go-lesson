package main

import (
    "fmt"
    "log"
    "os"
    "io"
)

func loggingSettings(logFile string) {
    logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    multiLogFile := io.MultiWriter(os.Stdout, logfile)
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log.SetOutput(multiLogFile)
}

func main() {
    _, err := os.Open("fdhgknad")
    if err != nil{
        log.Fatalln("exit", err)
    }

    log.Println("logging!")
    log.Printf("%T %v", "test", "test")

    log.Fatalf("%T %v", "test", "test")
    log.Fatalln("error")
    fmt.Println("ok!")
}