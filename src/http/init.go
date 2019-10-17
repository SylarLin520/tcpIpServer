package http

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "thrift"
    "fmt"
    "sync"
)

func Init() {
    router := gin.Default()
    router.GET("/someGet", getting)
    router.POST("/somePost", posting)
    router.POST("/2thrift", httpPost2Thrift)
    router.Run(":8080")
}

func getting(gin *gin.Context) {
    thrift.ClientTest("get")
    gin.String(http.StatusOK, "aaabbbccc")
}

func posting(gin *gin.Context) {
    req := &postRequest{}
    err := gin.ShouldBindJSON(req)
    if err != nil {
        fmt.Println("err: ", err.Error())
    }

    thrift.ClientTest(req.Name)
    fmt.Println(req)
}

type postRequest struct {
    Id   int64
    Name string
}

func RunHttpServer() {
    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("do http Init")
        Init()
    }()
    wg.Wait()
}

func httpPost2Thrift(gin *gin.Context){
    req := &postRequest{}
    err := gin.ShouldBindJSON(req)
    if err != nil {
        fmt.Println("err: ", err.Error())
    }

    thrift.ClientTest(req.Name)
    fmt.Println(req)
}