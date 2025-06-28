package client

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/Erik142/veil-configs/pkg/proto"
	"github.com/sirupsen/logrus"
)

// Run fetches the Nebula configuration from the server and saves it to a file.
func Run(client pb.NebulaConfigServiceClient, clientID, fileName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.GetNebulaConfig(ctx, &pb.GetNebulaConfigRequest{ClientId: clientID})
	if err != nil {
		return fmt.Errorf("could not get nebula config: %v", err)
	}
	logrus.Printf("Received Nebula Config for client %s", clientID)

	err = os.WriteFile(fileName, []byte(r.GetConfigContent()), 0644)
	if err != nil {
		return fmt.Errorf("failed to save config to file: %v", err)
	}
	logrus.Printf("Nebula configuration saved to %s", fileName)
	return nil
}
