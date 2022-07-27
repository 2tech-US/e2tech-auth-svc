package main

import (
	"fmt"
	"log"
	"net"

	"github.com/lntvan166/e2tech-auth-svc/internal/config"
	"github.com/lntvan166/e2tech-auth-svc/internal/db"
	"github.com/lntvan166/e2tech-auth-svc/internal/passenger"
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
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	passengerSvc := &passenger.ServiceClient{
		PassengerClient: passenger.InitServiceClient(&c),
	}

	s := services.Server{
		DB:           DB,
		PassengerSvc: passengerSvc,
		Jwt:          jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
