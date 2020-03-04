/**
*@program: GRPC-Study
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-04 14:51
 */
package main

import (
	"GRPC-Study/stream/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type ser struct {
}
// https://segmentfault.com/a/1190000016503114
func main() {
	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server, &ser{})

	listen, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Fatalln(err)
	}
	server.Serve(listen)
}

func (s *ser) List(req *pb.StreamRequest, stream pb.StreamService_ListServer) error {
	for n := 0; n < 6; n++ {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  req.Pt.Name,
				Value: req.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
	// 发送完成后 io.EOF
}

func (s *ser) Record(stream pb.StreamService_RecordServer) error {
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			// 如果关闭了 我们也关闭这个连接
			return stream.SendAndClose(&pb.StreamResponse{Pt:&pb.StreamPoint{Name:"aa",Value:123}})
		}
		if err != nil {
			return err
		}
		log.Printf("stream.Recv pt.name: %s, pt.value: %d", recv.Pt.Name, recv.Pt.Value)
	}
	return nil
}

func (s *ser) Route(stream pb.StreamService_RouteServer) error {
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		err = stream.Send(&pb.StreamResponse{Pt: &pb.StreamPoint{
			Name:  recv.Pt.Name,
			Value: recv.Pt.Value + 1,
		}})
		if err != nil {
			return err
		}
	}
	return nil
}
