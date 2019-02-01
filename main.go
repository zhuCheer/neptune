package main

import (
	"fmt"
	pb "github.com/zhuCheer/neptune/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"net"
	"flag"
	"log"
	"github.com/zhuCheer/neptune/service"
)

func startGrpc(host string){
	log.Println("bind host:"+ host)
	lis, err := net.Listen("tcp", host)
	if err != nil {
		fmt.Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server := &service.RedisSrv{}
	pb.RegisterNeptunePoolServer(s, server)

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Errorf("failed to serve: %v", err)
	}
}



func main(){
	var host string
	flag.StringVar(&host, "host", "0.0.0.0:50033", "input bind host")
	flag.Parse()

	fmt.Println("start pool")
	startGrpc(host)


}
