package server

import (
	"context"
	"fmt"
	"net"

	"github.com/Erik142/veil-configs/pkg/config"
	pb "github.com/Erik142/veil-configs/pkg/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// server is used to implement nebulaconfig.NebulaConfigServiceServer.
type GrpcServer struct {
	pb.UnimplementedNebulaConfigServiceServer
	ConfigStore config.ConfigStore
}

// GetNebulaConfig implements nebulaconfig.NebulaConfigServiceServer.
func (s *GrpcServer) GetNebulaConfig(ctx context.Context, in *pb.GetNebulaConfigRequest) (*pb.GetNebulaConfigResponse, error) {
	logrus.Printf("Received request for client ID: %s", in.GetClientId())
	configContent, err := s.ConfigStore.GetConfig(in.GetClientId())
	if err != nil {
		return nil, err
	}
	return &pb.GetNebulaConfigResponse{ConfigContent: configContent}, nil
}

// StartServer starts the gRPC server.
func StartServer(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNebulaConfigServiceServer(s, &GrpcServer{ConfigStore: config.NewInMemoryConfigStore()})
	logrus.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}