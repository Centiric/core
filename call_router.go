package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	corepb "github.com/Centiric/core/proto/core"
	mediapb "github.com/Centiric/core/proto/media"
)

func (s *voipServer) RouteCall(ctx context.Context, req *corepb.CallRequest) (*corepb.CallResponse, error) {
	// Artık zerolog ile yapılandırılmış loglama yapıyoruz
	log.Info().Str("from", req.From).Str("to", req.To).Msg("RouteCall isteği alındı")

	mediaAddr := s.config.Services.Media.Address
	log.Debug().Str("address", mediaAddr).Msg("Media servisine bağlanılıyor...")

	conn, err := grpc.NewClient(mediaAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Err(err).Str("address", mediaAddr).Msg("Media servisine bağlanılamadı")
		return &corepb.CallResponse{Status: corepb.CallResponse_FAILED}, nil
	}
	defer conn.Close()

	mediaClient := mediapb.NewMediaManagerClient(conn)

	mediaCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	log.Debug().Msg("Media servisinden RTP portu isteniyor...")
	mediaRes, err := mediaClient.AllocatePort(mediaCtx, &mediapb.AllocatePortRequest{})
	if err != nil {
		log.Error().Err(err).Msg("Media servisinden port alınamadı")
		return &corepb.CallResponse{Status: corepb.CallResponse_FAILED}, nil
	}

	rtpPort := mediaRes.GetPort()
	log.Info().Uint32("rtp_port", rtpPort).Msg("Media servisinden port başarıyla alındı")

	sessionID := "sess_" + time.Now().Format("20060102150405")

	return &corepb.CallResponse{
		Status:    corepb.CallResponse_OK,
		SessionId: sessionID,
		RtpPort:   rtpPort, // <-- TEMİZ VE DOĞRU YÖNTEM
	}, nil
}
