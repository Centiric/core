package main

import (
	"context"
	"fmt"
	"time"

	// Yerel proto paketimizi import ediyoruz
	"github.com/Centiric/core/proto"
)

// Dikkat: Fonksiyon imzalarındaki tiplerin başına 'proto.' ekledik
func (s *voipServer) RouteCall(ctx context.Context, req *proto.CallRequest) (*proto.CallResponse, error) {
	if req.To == "" {
		return &proto.CallResponse{Status: proto.CallResponse_FAILED}, nil
	}
	return &proto.CallResponse{
		Status:    proto.CallResponse_OK,
		SessionId: fmt.Sprintf("sess_%x", time.Now().UnixNano()),
	}, nil
}
