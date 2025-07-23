package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/Centiric/core/proto"
)

func (s *voipServer) RouteCall(ctx context.Context, req *pb.CallRequest) (*pb.CallResponse, error) {
	if req.To == "" {
		return &pb.CallResponse{Status: pb.CallResponse_FAILED}, nil
	}
	return &pb.CallResponse{
		Status:    pb.CallResponse_OK,
		SessionId: fmt.Sprintf("sess_%x", time.Now().UnixNano()),
	}, nil
}
