// C:\centric\core\main.go

package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// 'proto' yerine 'corepb' takma adını kullanıyoruz
	corepb "github.com/Centiric/core/proto/core"
)

// 'proto.' yerine 'corepb.' kullanıyoruz
type voipServer struct {
	corepb.UnimplementedVoipCoreServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// 'proto.' yerine 'corepb.' kullanıyoruz
	corepb.RegisterVoipCoreServer(s, &voipServer{})

	reflection.Register(s)

	log.Printf("Server started at :%s", "50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
