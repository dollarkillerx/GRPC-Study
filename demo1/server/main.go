/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-10
* Time: 下午2:52
* */
package main

import (
	"GRPC-Study/demo1/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type server struct {}

func main() {
	// 创建一个tcp服务
	listener, e := net.Listen("tcp", ":9001")
	if e != nil {
		panic(e.Error())
	}

	// 创建一个没有注册的grpc
	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv,&server{})// 注册上去
	reflection.Register(srv) // 注册在给定的gRPC服务器上注册服务器反射服务

	// 监听
	if e := srv.Serve(listener);e != nil {
		panic(e.Error())
	}
}


// 要去实现service.proto定义的方法
func (s *server) Add(ctx context.Context,req *proto.Request) (*proto.Response,error) {
	a,b := req.GetA(),req.GetB()
	result := a + b
	return &proto.Response{Result:result},nil
}

func (s *server) Multiply(ctx context.Context,req *proto.Request) (*proto.Response,error) {
	b,a := req.GetB(),req.GetA()
	result := b*a
	return &proto.Response{
		Result:result,
	},nil
}