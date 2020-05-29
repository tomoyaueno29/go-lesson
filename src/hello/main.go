package main

import (
    
)

type ConfigList struct {
    Port      int
    DbName    string
    SQLDriver string
}

var Config ConfigList

func init() {
    cfg, _ := inim.Load("config.ini")
    
}

func main() {
    

}