package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	rn "rendering-engine/packages/random-number"
)

var PORT = ":9001"

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("error starting grpc server %v\n", err)
	}
	s := rn.Server{}
	grpcServer := grpc.NewServer()

	rn.RegisterRenderingEngineServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("error starting grpc server %v\n", err)
	}
}
