package thrift

import (
    "fmt"
    "thriftExample/gen-go/echo"
    "github.com/apache/thrift/lib/go/thrift"
    "context"
)

func Init() {
    fmt.Println("[thrift] Init start...")
    transport, err := thrift.NewTServerSocket(":9898")
    if err != nil {
        panic(err)
    }

    handler := &EchoServer{}
    processor := echo.NewEchoProcessor(handler)

    transportFactory := thrift.NewTBufferedTransportFactory(8192)
    protocolFactory := thrift.NewTCompactProtocolFactory()
    server := thrift.NewTSimpleServer4(
        processor,
        transport,
        transportFactory,
        protocolFactory,
    )

    if err := server.Serve(); err != nil {
        panic(err)
    }
    fmt.Println("thrift Init end...")

}

type EchoServer struct {
}

func (e *EchoServer) Echo(ctx context.Context, req *echo.EchoReq) (*echo.EchoRes, error) {
    fmt.Printf("message from client: %v\n", req.GetMsg())

    res := &echo.EchoRes{
        Msg: "success",
    }

    return res, nil
}
