/**
*@program: GRPC-Study
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-04 14:52
 */
package main

import (
	"GRPC-Study/stream/pb"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := pb.NewStreamServiceClient(conn)

	//if err := printLists(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "Stream Lists...", Value: 18650}}); err != nil {
	//	log.Fatalln(err)
	//}

	//if err := printRecord(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "Stream Lists...", Value: 18651}}); err != nil {
	//	log.Fatalln(err)
	//}

	if err := printRoute(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "Stream Route...", Value: 18652}}); err != nil {
		log.Fatalln(err)
	}
}

func printLists(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	list, err := client.List(context.TODO(), r)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		recv, err := list.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp name: %s value: %d", recv.Pt.Name, recv.Pt.Value)
	}
	return nil
}

func printRecord(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	record, err := client.Record(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}
	for i:=0;i<20;i++ {
		r.Pt.Value += int32(i)
		err := record.Send(r)
		if err != nil {
			return err
		}
	}
	// 客户端流 客户端负责关闭
	resp, err := record.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	return nil
}

func printRoute(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	route, err := client.Route(context.TODO())
	if err != nil {
		return err
	}
	
	for i:=0;i<20;i++ {
		err := route.Send(r)
		if err != nil {
			return err
		}
		recv, err := route.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		recv.Pt = r.Pt
		log.Printf("resp: pj.name: %s, pt.value: %d", recv.Pt.Name, recv.Pt.Value)
	}
	route.CloseSend()
	return nil
}
