package main

import (
	"log"
	"net"

	"github.com/Justyer/bj21"
	"github.com/Justyer/bj21/pb"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("lis_err: %s", err)
	}
	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	pb.RegisterBJ21Server(server, &bj21.BJ21Impl{})
	server.Serve(lis)
}
