package client

import (
	"context"
	"fmt"

	"github.com/lntvan166/e2tech-auth-svc/internal/config"
	"github.com/lntvan166/e2tech-auth-svc/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DriverServiceClient struct {
	Client pb.DriverServiceClient
}

func InitDriverServiceClient(c *config.Config) pb.DriverServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.DriverSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewDriverServiceClient(cc)
}

type CreateDriverRequest struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

func (s *DriverServiceClient) CreateDriver(ctx context.Context, req *CreateDriverRequest) (*pb.CreateDriverResponse, error) {
	return s.Client.CreateDriver(ctx, &pb.CreateDriverRequest{
		Phone: req.Phone,
		Name:  req.Name,
	})

}
