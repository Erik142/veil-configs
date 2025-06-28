package app

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	client "github.com/Erik142/veil-configs/internal/client"
	pb "github.com/Erik142/veil-configs/pkg/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

// MockNebulaConfigServiceClient is a mock implementation of the gRPC client for testing.
type MockNebulaConfigServiceClient struct {
	pb.UnimplementedNebulaConfigServiceServer
	ConfigContent string
	Error         error
}

func (m *MockNebulaConfigServiceClient) GetNebulaConfig(ctx context.Context, in *pb.GetNebulaConfigRequest, opts ...grpc.CallOption) (*pb.GetNebulaConfigResponse, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	return &pb.GetNebulaConfigResponse{ConfigContent: m.ConfigContent}, nil
}

func TestFetchAndSaveConfig(t *testing.T) {
	// Create a mock client
	mockClient := &MockNebulaConfigServiceClient{
		ConfigContent: "test_config_content",
		Error:         nil,
	}

	// Create a temporary file for saving the config
	tmpFile, err := ioutil.TempFile("", "test_config_*.yaml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	// Call the function to test
	err = client.Run(mockClient, "test_client", tmpFile.Name())
	assert.NoError(t, err)

	// Verify the content of the saved file
	content, err := ioutil.ReadFile(tmpFile.Name())
	assert.NoError(t, err)
	assert.Equal(t, "test_config_content", string(content))

	// Test error case
	mockClient.Error = assert.AnError // Set a mock error
		err = client.Run(mockClient, "test_client", tmpFile.Name())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "could not get nebula config")
}
