package client

import (
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

type CallCenterServiceClient struct {
	Client pb.CallCenterServiceClient
}

func InitServiceClient(c *config.Config) pb.CallCenterServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.CallCenterSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCallCenterServiceClient(cc)
}
