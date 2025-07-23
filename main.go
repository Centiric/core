package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	// YENİ: Reflection paketini import ediyoruz
	"github.com/Centiric/core/proto"
	"google.golang.org/grpc/reflection"
)

type voipServer struct {
	proto.UnimplementedVoipCoreServer
}

func main() {
	lis, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterVoipCoreServer(s, &voipServer{})

	// YENİ: Bu satır, gRPC sunucusuna yansıma özelliğini kaydediyor.
	// Artık grpcurl gibi araçlar sunucuyu keşfedebilir.
	reflection.Register(s)

	log.Printf("Server started at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
