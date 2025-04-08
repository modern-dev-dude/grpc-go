package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	rn "rendering-engine/packages/random-number"
)

var PORT = ":9001"

/**
* The intent of this is to simulate some large workload in a separate server
* then integrate into the client
 */
func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("error starting grpc server %v\n", err)
	}
	s := rn.Server{}
	grpcServer := grpc.NewServer()

	rn.RegisterRandomNumberServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("error starting grpc server %v\n", err)
	}
}
