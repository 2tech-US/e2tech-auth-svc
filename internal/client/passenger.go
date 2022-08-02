package client

import (
	"context"

	"github.com/lntvan166/e2tech-auth-svc/internal/pb"
)

type CreatePassengerRequest struct {
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
}

func (s *PassengerServiceClient) CreatePassenger(ctx context.Context, req *CreatePassengerRequest) (*pb.CreatePassengerResponse, error) {
	return s.Client.CreatePassenger(ctx, &pb.CreatePassengerRequest{
		Phone:    req.Phone,
		Password: req.Password,
		Name:     req.Name,
	})

}

type VerifyPassengerRequest struct {
	Phone string `json:"phone"`
}

func (s *PassengerServiceClient) VerifyPassenger(ctx context.Context, req *pb.VerifyPassengerRequest) (*pb.VerifyPassengerResponse, error) {
	return s.Client.VerifyPassenger(ctx, &pb.VerifyPassengerRequest{
		Phone: req.Phone,
	})
}
