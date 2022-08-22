package main

import (
	"fmt"
	"log"
	"net"

	"github.com/lntvan166/e2tech-auth-svc/internal/client"
	"github.com/lntvan166/e2tech-auth-svc/internal/config"
	"github.com/lntvan166/e2tech-auth-svc/internal/db"
	"github.com/lntvan166/e2tech-auth-svc/internal/pb"
	"github.com/lntvan166/e2tech-auth-svc/internal/services"
	"github.com/lntvan166/e2tech-auth-svc/internal/utils"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	DB := db.Connect(c.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	passengerSvc := &client.PassengerServiceClient{
		Client: client.InitPassengerServiceClient(&c),
	}
	driverSvc := &client.DriverServiceClient{
		Client: client.InitDriverServiceClient(&c),
	}
	callcenterSvc := &client.CallCenterServiceClient{
		Client: client.InitCallCenterServiceClient(&c),
	}

	s := services.Server{
		DB:            DB,
		PassengerSvc:  passengerSvc,
		DriverSvc:     driverSvc,
		CallCenterSvc: callcenterSvc,
		Jwt:           jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
