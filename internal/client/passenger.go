package client

import (
	"context"
	"fmt"

	"github.com/lntvan166/e2tech-auth-svc/internal/config"
	"github.com/lntvan166/e2tech-auth-svc/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PassengerServiceClient struct {
	Client pb.PassengerServiceClient
}

func InitPassengerServiceClient(c *config.Config) pb.PassengerServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.PassengerSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewPassengerServiceClient(cc)
}

type CreatePassengerRequest struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

func (s *PassengerServiceClient) CreatePassenger(ctx context.Context, req *CreatePassengerRequest) (*pb.CreatePassengerResponse, error) {
	return s.Client.CreatePassenger(ctx, &pb.CreatePassengerRequest{
		Phone: req.Phone,
		Name:  req.Name,
	})

}
