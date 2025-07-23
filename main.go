package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
	"google.golang.org/grpc"
	pb "github.com/Centiric/core/proto"
)

type voipServer struct {
	pb.UnimplementedVoipCoreServer
}

func main() {
	lis, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterVoipCoreServer(s, &voipServer{})
	log.Printf("Server started at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}