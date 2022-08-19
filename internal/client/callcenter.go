package client

import (
	"context"

	"github.com/lntvan166/e2tech-auth-svc/internal/pb"
)

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
