package main

import (
	"log"
	"net"

	"github.com/Centiric/core/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type voipServer struct {
	proto.UnimplementedVoipCoreServer
}

func main() {
	// DİKKAT: Sadece IPv4 loopback adresini dinliyoruz.
	address := ":50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterVoipCoreServer(s, &voipServer{})
	reflection.Register(s)

	log.Printf("Server started at %s", address) // Log mesajını da güncelledik.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
