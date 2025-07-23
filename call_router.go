// call_router.go
package main

import (
	"context"
	pb "github.com/Centiric/core/proto"
)

func (s *voipServer) RouteCall(ctx context.Context, req *pb.CallRequest) (*pb.CallResponse, error) {
	if req.To == "" {
		return &pb.CallResponse{Status: pb.CallResponse_FAILED}, nil
	}
	return &pb.CallResponse{
		Status:    pb.CallResponse_OK,
		SessionId: "sess_" + fmt.Sprintf("%x", time.Now().UnixNano()),
	}, nil
}