/**
 * @Author: DollarKillerX
 * @Description: service.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午9:51 2020/1/14
 */
package main

import (
	proto "GRPC-Study/demo2/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func main() {
	listener, e := net.Listen("tcp", ":9901")
	if e != nil {
		log.Fatalln(e)
	}

	//credentials.NewServerTLSFromFile
	tls, e := credentials.NewServerTLSFromFile("/home/dollarkiller/Github/GRPC-Study/demo2/key/server.pem", "/home/dollarkiller/Github/GRPC-Study/demo2/key/server.key")
	if e != nil {
		log.Fatalln(e)
	}
	server := grpc.NewServer(grpc.Creds(tls))

	proto.RegisterHelloServer(server, &hello{})

	e = server.Serve(listener)
	if e != nil {
		log.Fatalln(e)
	}
}

type hello struct {
}

func (h *hello) Result(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	log.Println(request.Msg)
	return &proto.Response{Code: "200", Msg: "Hello GRPC"}, nil
}
