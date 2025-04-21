package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"rendering-engine/packages/renderer"
)

var PORT = ":9000"

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("error starting grpc server %v\n", err)
	}
	s := go_renderer.Server{}
	grpcServer := grpc.NewServer()

	go_renderer.RegisterRenderingEngineServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("error starting grpc server %v\n", err)
	}
}
