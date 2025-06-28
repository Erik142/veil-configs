package main

import (
	"context"
	"io/ioutil"
	"net"
	"os"
	"testing"

	pb "github.com/Erik142/veil-configs/pkg/proto"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

// MockNebulaConfigServiceServer is a mock implementation of the gRPC server for testing.
type MockNebulaConfigServiceServer struct {
	pb.UnimplementedNebulaConfigServiceServer
	ConfigContent string
	Error         error
}

func (m *MockNebulaConfigServiceServer) GetNebulaConfig(ctx context.Context, in *pb.GetNebulaConfigRequest) (*pb.GetNebulaConfigResponse, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	return &pb.GetNebulaConfigResponse{ConfigContent: m.ConfigContent}, nil
}

func TestFetchAndSaveConfig(t *testing.T) {
	// Create a buffer for the in-memory gRPC connection
	lis := bufconn.Listen(1024 * 1024)
	defer lis.Close()

	// Create a mock server
	mockServer := &MockNebulaConfigServiceServer{
		ConfigContent: "test_config_content",
		Error:         nil,
	}

	// Start the gRPC server with the mock service
	s := grpc.NewServer()
	pb.RegisterNebulaConfigServiceServer(s, mockServer)
	go func() {
		if err := s.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			logrus.Fatalf("Mock server exited with error: %v", err)
		}
	}()
	defer s.Stop()

	// Create a client connection to the mock server
	conn, err := grpc.DialContext(context.Background(), "bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	defer conn.Close()

	// Create a temporary file for saving the config
	tmpFile, err := ioutil.TempFile("", "test_config_*.yaml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	// Call the function to test
	err = FetchAndSaveConfig(conn, "test_client", tmpFile.Name())
	assert.NoError(t, err)

	// Verify the content of the saved file
	content, err := ioutil.ReadFile(tmpFile.Name())
	assert.NoError(t, err)
	assert.Equal(t, "test_config_content", string(content))

	// Test error case
	mockServer.Error = assert.AnError // Set a mock error
	err = FetchAndSaveConfig(conn, "test_client", tmpFile.Name())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "could not get nebula config")
}
