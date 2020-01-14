/**
 * @Author: DollarKillerX
 * @Description: main.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:22 2020/1/14
 */
package main

import (
	"GRPC-Study/demo3/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

func main() {
	tc, e := credentials.NewServerTLSFromFile("/home/dollarkiller/Github/GRPC-Study/demo3/key/server.pem", "/home/dollarkiller/Github/GRPC-Study/demo3/key/server.key")
	if e != nil {
		log.Fatalln(e)
	}

	listener, e := net.Listen("tcp", ":9901")
	if e != nil {
		log.Fatalln(e)
	}
	server := grpc.NewServer(grpc.Creds(tc))
	pb.RegisterHelloServer(server, &hello{})

	e = server.Serve(listener)
	if e != nil {
		log.Fatalln(e)
	}
}

type hello struct {
}

func (h *hello) Result(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.Resp{Code: "401", Msg: "No Token"}, nil
	}
	appid, ok := md["appid"]
	if ok {
		log.Println(appid)
	}

	appkey, ok := md["appkey"]
	if ok {
		log.Println(appkey)
	}

	return &pb.Resp{Code: "200", Msg: req.Msg + "   GRPC"}, nil
}
