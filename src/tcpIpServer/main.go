package main

import (
    h "http"
    "thrift"
    "fmt"
)

func main() {
    go h.RunHttpServer()
    fmt.Println("do thrift Init")
    thrift.Init()
}
