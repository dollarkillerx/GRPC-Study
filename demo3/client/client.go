/**
 * @Author: DollarKillerX
 * @Description: client.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:20 2020/1/14
 */
package main

import (
	"GRPC-Study/demo3/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	var opts []grpc.DialOption
	if OpenTls {
		tc, e := credentials.NewClientTLSFromFile("/home/dollarkiller/Github/GRPC-Study/demo3/key/server.pem", "ppx")
		if e != nil {
			log.Fatalln(e)
		}
		opts = append(opts, grpc.WithTransportCredentials(tc))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// 使用自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(&customCredential{}))

	conn, e := grpc.Dial(":9901", opts...)
	if e != nil {
		log.Fatalln(e)
	}
	client := pb.NewHelloClient(conn)
	resp, e := client.Result(context.TODO(), &pb.Req{Msg: "Hello World"})
	if e != nil {
		log.Fatalln(e)
	}
	log.Println(resp)
}

var OpenTls = true

// customCredential自定义认证
type customCredential struct{}

func (c *customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "0001",
		"appkey": "key",
	}, nil
}

func (c *customCredential) RequireTransportSecurity() bool {
	if OpenTls {
		return true
	}
	return false
}
