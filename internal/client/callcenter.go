package client

import (
	"context"
	"fmt"

	"github.com/lntvan166/e2tech-auth-svc/internal/config"
	"github.com/lntvan166/e2tech-auth-svc/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CallCenterServiceClient struct {
	Client pb.CallCenterServiceClient
}

func InitCallCenterServiceClient(c *config.Config) pb.CallCenterServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.CallCenterSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCallCenterServiceClient(cc)
}

type CreateEmployeeRequest struct {
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

func (s *CallCenterServiceClient) CreateEmployee(ctx context.Context, req *CreateEmployeeRequest) (*pb.GetEmployeeResponse, error) {
	return s.Client.CreateEmployee(ctx, &pb.CreateCallCenterEmployeeRequest{
		Phone: req.Phone,
		Role:  req.Role,
	})

}
