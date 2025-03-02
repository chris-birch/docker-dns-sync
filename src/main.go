package main

import (
	"github.com/chris-birch/docker-dns-sync/proto/technitium/v1/service"
	"github.com/chris-birch/docker-dns-sync/src/technitium"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	// create a TCP listener on the specified port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("tcp connection failed: %v", err)
	}
	log.Printf("listening at %v", lis.Addr())

	// create a gRPC server instance
	s := grpc.NewServer()
	reflection.Register(s)

	// Assign the GRPc services
	technitiumService := technitium.Service{}

	// Register the services
	service.RegisterTechnitiumServiceServer(s, &technitiumService)

	// start listening to requests
	if err := s.Serve(lis); err != nil {
		log.Fatalf("grpc server failed: %v", err)
	}
}
