/**
 * @Author: DollarKillerX
 * @Description: client.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午9:49 2020/1/14
 */
package main

import (
	proto "GRPC-Study/demo2/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	tc, err := credentials.NewClientTLSFromFile("/home/dollarkiller/Github/GRPC-Study/demo2/key/server.pem", "ssr")
	if err != nil {
		log.Fatalln(err)
	}

	conn, e := grpc.Dial(":9901", grpc.WithTransportCredentials(tc))
	if e != nil {
		panic(e)
	}
	defer conn.Close()

	client := proto.NewHelloClient(conn)
	response, e := client.Result(context.TODO(), &proto.Request{Msg: "Hello World"})
	if e != nil {
		panic(e)
	}
	log.Println(response)
}
