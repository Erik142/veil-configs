package app

import (
	"context"
	"net"
	"testing"

	internal_server "github.com/Erik142/veil-configs/internal/server"
	"github.com/Erik142/veil-configs/pkg/config"
	pb "github.com/Erik142/veil-configs/pkg/proto"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	go func() {
		s := grpc.NewServer()
		pb.RegisterNebulaConfigServiceServer(s, &internal_server.GrpcServer{ConfigStore: config.NewInMemoryConfigStore()})
		if err := s.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			logrus.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGetNebulaConfig(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	defer conn.Close()

	client := pb.NewNebulaConfigServiceClient(conn)

	// Test with existing client ID
	req := &pb.GetNebulaConfigRequest{ClientId: "client1"}
	res, err := client.GetNebulaConfig(ctx, req)
	assert.NoError(t, err)
	assert.Contains(t, res.GetConfigContent(), "client1_ca_cert_content")

	// Test with non-existing client ID
	req = &pb.GetNebulaConfigRequest{ClientId: "nonexistent_client"}
	res, err = client.GetNebulaConfig(ctx, req)
	assert.Error(t, err)
	assert.Nil(t, res)
}
