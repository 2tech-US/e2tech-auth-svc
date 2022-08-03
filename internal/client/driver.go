package client

import (
	"context"

	"github.com/lntvan166/e2tech-auth-svc/internal/pb"
)

type CreateDriverRequest struct {
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
}

func (s *DriverServiceClient) CreateDriver(ctx context.Context, req *CreateDriverRequest) (*pb.CreateDriverResponse, error) {
	return s.Client.CreateDriver(ctx, &pb.CreateDriverRequest{
		Phone:    req.Phone,
		Password: req.Password,
		Name:     req.Name,
	})

}

type VerifyDriverRequest struct {
	Phone string `json:"phone"`
}

func (s *DriverServiceClient) VerifyDriver(ctx context.Context, req *pb.VerifyDriverRequest) (*pb.VerifyDriverResponse, error) {
	return s.Client.VerifyDriver(ctx, &pb.VerifyDriverRequest{
		Phone: req.Phone,
	})
}
