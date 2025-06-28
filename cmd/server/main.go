package main

import (
	"context"
	"log"
	"net"

	"github.com/Erik142/veil-configs/pkg/config"
	pb "github.com/Erik142/veil-configs/pkg/proto"
	"google.golang.org/grpc"
)

// server is used to implement nebulaconfig.NebulaConfigServiceServer.
type server struct {
	pb.UnimplementedNebulaConfigServiceServer
	configStore config.ConfigStore
}

// GetNebulaConfig implements nebulaconfig.NebulaConfigServiceServer.
func (s *server) GetNebulaConfig(ctx context.Context, in *pb.GetNebulaConfigRequest) (*pb.GetNebulaConfigResponse, error) {
	log.Printf("Received request for client ID: %s", in.GetClientId())
	configContent, err := s.configStore.GetConfig(in.GetClientId())
	if err != nil {
		return nil, err
	}
	return &pb.GetNebulaConfigResponse{ConfigContent: configContent}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNebulaConfigServiceServer(s, &server{configStore: config.NewInMemoryConfigStore()})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}