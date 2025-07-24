package main

import (
	"fmt"
	"net"

	"github.com/Centiric/core/config"
	"github.com/Centiric/core/logger"
	"github.com/rs/zerolog/log"

	corepb "github.com/Centiric/core/proto/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type voipServer struct {
	corepb.UnimplementedVoipCoreServer
	config config.Config // Konfigürasyonu sunucu içinde tutalım
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Konfigürasyon yüklenemedi")
	}

	logger.Initialize(cfg.Log.Level)

	listenAddr := fmt.Sprintf("%s:%d", cfg.Grpc.Host, cfg.Grpc.Port)
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal().Err(err).Str("address", listenAddr).Msg("Dinleme başlatılamadı")
	}

	s := grpc.NewServer()
	corepb.RegisterVoipCoreServer(s, &voipServer{config: cfg})
	reflection.Register(s)

	log.Info().Str("address", listenAddr).Msg("gRPC sunucusu başlatıldı")
	if err := s.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("Sunucu çalışırken hata oluştu")
	}
}
