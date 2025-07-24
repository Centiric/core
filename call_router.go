// C:\centric\core\call_router.go

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	corepb "github.com/Centiric/core/proto/core"
	mediapb "github.com/Centiric/core/proto/media"
)

func (s *voipServer) RouteCall(ctx context.Context, req *corepb.CallRequest) (*corepb.CallResponse, error) {
	log.Printf("[CORE] RouteCall isteği alındı: From=%s, To=%s", req.From, req.To)

	log.Println("[CORE] Media servisine bağlanılıyor (localhost:50052)...")

	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// ...
	}
	defer conn.Close()

	mediaClient := mediapb.NewMediaManagerClient(conn)

	log.Println("[CORE] Media servisinden RTP portu isteniyor...")

	// --- DÜZELTME BURADA ---
	// `context.Background()` yerine, fonksiyona gelen orijinal `ctx`'i kullanıyoruz.
	// Bu, eğer `signal`'dan gelen isteğin bir timeout'u varsa, o timeout'un
	// `media`'ya yapılan isteğe de yansımasını sağlar.
	mediaCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	// ----------------------
	defer cancel()

	// mediaCtx'i kullanarak isteği gönder
	mediaRes, err := mediaClient.AllocatePort(mediaCtx, &mediapb.AllocatePortRequest{})
	if err != nil {
		log.Printf("[HATA] Media servisinden port alınamadı: %v", err)
		return &corepb.CallResponse{Status: corepb.CallResponse_FAILED}, nil
	}

	rtpPort := mediaRes.GetPort()
	log.Printf("[CORE] Media servisinden port başarıyla alındı: %d", rtpPort)

	sessionID := fmt.Sprintf("sess_%x_port_%d", time.Now().UnixNano(), rtpPort)

	return &corepb.CallResponse{
		Status:    corepb.CallResponse_OK,
		SessionId: sessionID,
	}, nil
}
