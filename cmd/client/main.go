package main

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/Erik142/veil-configs/pkg/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func FetchAndSaveConfig(conn grpc.ClientConnInterface, clientID, fileName string) error {
	c := pb.NewNebulaConfigServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetNebulaConfig(ctx, &pb.GetNebulaConfigRequest{ClientId: clientID})
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

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	clientID := "client1"
	fileName := fmt.Sprintf("nebula_config_%s.yaml", clientID)

	if err := FetchAndSaveConfig(conn, clientID, fileName); err != nil {
		logrus.Fatalf("failed to fetch and save config: %v", err)
	}
}
