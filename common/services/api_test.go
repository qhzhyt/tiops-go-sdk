package services

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"math"
	"net"
	"testing"
	"tiops/common/models"
)

type APIServer struct {
	count int
}

func (s *APIServer) PostLog(ctx context.Context, req *models.Log) (*APIResponse, error) {

	s.count ++

	if s.count % 1000 == 999{
		log.Println(s.count)
	}

	return &APIResponse{Code: 123456, Message: req.Content}, nil
}

func TestAPIServiceServer(t *testing.T) {
	sock, err := net.Listen("tcp", ":5555")
	if err != nil {
		t.Error(err)
		return
	}

	var options = []grpc.ServerOption{
		grpc.MaxRecvMsgSize(math.MaxInt32),
		grpc.MaxSendMsgSize(1073741824),
	}
	s := grpc.NewServer(options...)

	myServer := &APIServer{}
	RegisterAPIServiceServer(s, myServer)
	if err := s.Serve(sock); err != nil {
		t.Fatalf("failed to serve: %v", err)
	}
}
