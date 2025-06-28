package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/Erik142/veil-configs/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewNebulaConfigServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	clientID := "client1" // You can change this to client2 or any other ID
	r, err := c.GetNebulaConfig(ctx, &pb.GetNebulaConfigRequest{ClientId: clientID})
	if err != nil {
		log.Fatalf("could not get nebula config: %v", err)
	}
	log.Printf("Received Nebula Config for client %s", clientID)

	// Save the configuration to a file
	fileName := fmt.Sprintf("nebula_config_%s.yaml", clientID)
	err = os.WriteFile(fileName, []byte(r.GetConfigContent()), 0644)
	if err != nil {
		log.Fatalf("failed to save config to file: %v", err)
	}
	log.Printf("Nebula configuration saved to %s", fileName)
}
