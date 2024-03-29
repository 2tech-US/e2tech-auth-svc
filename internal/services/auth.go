package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	client "github.com/lntvan166/e2tech-auth-svc/internal/client"
	"github.com/lntvan166/e2tech-auth-svc/internal/db"
	"github.com/lntvan166/e2tech-auth-svc/internal/pb"
	"github.com/lntvan166/e2tech-auth-svc/internal/utils"
)

type Server struct {
	DB            *db.Queries
	Jwt           utils.JwtWrapper
	PassengerSvc  *client.PassengerServiceClient
	DriverSvc     *client.DriverServiceClient
	CallCenterSvc *client.CallCenterServiceClient
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	_, err := s.DB.GetUserByPhone(ctx, req.Phone)
	if err != sql.ErrNoRows {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "auth: user already exists",
		}, nil
	}

	if req.Role == utils.PASSENGER {
		rsp, err := s.PassengerSvc.CreatePassenger(ctx, &client.CreatePassengerRequest{
			Phone: req.Phone,
			Name:  req.Name,
		})
		if err != nil {
			return &pb.RegisterResponse{
				Status: http.StatusBadGateway,
				Error:  fmt.Sprintf("passenger service error: %s", err),
			}, nil
		}
		if rsp.Status != http.StatusCreated {
			return &pb.RegisterResponse{
				Status: rsp.Status,
				Error:  rsp.Error,
			}, nil
		}
	}
	if req.Role == utils.DRIVER {
		rsp, err := s.DriverSvc.CreateDriver(ctx, &client.CreateDriverRequest{
			Phone: req.Phone,
			Name:  req.Name,
		})
		if err != nil {
			return &pb.RegisterResponse{
				Status: http.StatusBadGateway,
				Error:  fmt.Sprintf("driver service error: %s", err),
			}, nil
		}
		if rsp.Status != http.StatusCreated {
			return &pb.RegisterResponse{
				Status: rsp.Status,
				Error:  rsp.Error,
			}, nil
		}
	}

	if req.Role == utils.CALLCENTER_CREATOR || req.Role == utils.CALLCENTER_LOCATOR || req.Role == utils.CALLCENTER_MANAGER {
		rsp, err := s.CallCenterSvc.CreateEmployee(ctx, &client.CreateEmployeeRequest{
			Phone: req.Phone,
			Role:  req.Role,
		})
		if err != nil {
			return &pb.RegisterResponse{
				Status: http.StatusBadGateway,
				Error:  fmt.Sprintf("Callcenter service error: %s", err),
			}, nil
		}
		if rsp.Status != http.StatusCreated {
			return &pb.RegisterResponse{
				Status: rsp.Status,
				Error:  rsp.Error,
			}, nil
		}
	}

	arg := db.CreateUserParams{
		Phone:    req.Phone,
		Password: utils.HashPassword(req.Password),
		Role:     req.Role,
	}

	_, err = s.DB.CreateUser(ctx, arg)
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("auth: create user error: %s", err),
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.DB.GetUserByPhone(ctx, req.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.LoginResponse{
				Status: http.StatusBadRequest,
				Error:  "auth: user not found",
			}, nil
		}
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("auth: get user error: %s", err),
		}, nil
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return &pb.LoginResponse{
			Status: http.StatusUnauthorized,
			Error:  "invalid password",
		}, nil
	}

	// if user.Role != req.Role {
	// 	return &pb.LoginResponse{
	// 		Status: http.StatusUnauthorized,
	// 		Error:  fmt.Sprintf("you are %s not %s", user.Role, req.Role),
	// 	}, nil
	// }

	_, err = s.DB.UpdateDeviceToken(ctx, db.UpdateDeviceTokenParams{
		Phone:       req.Phone,
		DeviceToken: utils.NullString(req.DeviceToken),
	})
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("auth: update device token error: %s", err),
		}, nil
	}

	token, err := s.Jwt.GenerateToken(user)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("generate token error: %s", err),
		}, nil
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	token := req.Token
	claims, err := s.Jwt.ValidateToken(token)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusUnauthorized,
			Error:  fmt.Sprintf("auth token error: %s", err),
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		Phone:  claims.Phone,
		Role:   claims.Role,
	}, nil
}

func (s *Server) Verify(ctx context.Context, req *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	token := req.Token
	claims, err := s.Jwt.ValidateToken(token)
	if err != nil {
		return &pb.VerifyResponse{
			Status: http.StatusUnauthorized,
			Error:  fmt.Sprintf("auth token error: %s", err),
		}, nil
	}

	if claims.Role != utils.ADMIN {
		return &pb.VerifyResponse{
			Status: http.StatusUnauthorized,
			Error:  "only admin can verify",
		}, nil
	}

	phone := req.Phone
	user, err := s.DB.GetUserByPhone(ctx, phone)
	if err != nil {
		return &pb.VerifyResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("user not found: %s", err),
		}, nil
	}

	if user.Verified {
		return &pb.VerifyResponse{
			Status: http.StatusBadRequest,
			Error:  "user already verified",
		}, nil
	}

	user.Verified = true
	_, err = s.DB.Verify(ctx, phone)
	if err != nil {
		return &pb.VerifyResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("verify error: %s", err),
		}, nil
	}

	return &pb.VerifyResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) GetDeviceToken(ctx context.Context, req *pb.GetDeviceTokenRequest) (*pb.GetDeviceTokenResponse, error) {
	phone := req.Phone
	user, err := s.DB.GetUserByPhone(ctx, phone)
	if err != nil {
		return &pb.GetDeviceTokenResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("user not found: %s", err),
		}, nil
	}

	return &pb.GetDeviceTokenResponse{
		Status: http.StatusOK,
		Token:  user.DeviceToken.String,
	}, nil
}
