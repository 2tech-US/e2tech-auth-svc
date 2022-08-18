package client

import (
	"context"

	"github.com/lntvan166/e2tech-auth-svc/internal/pb"
)

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
