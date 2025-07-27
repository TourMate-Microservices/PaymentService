package user

import (
	"context"
	"log"
	"os"
	"tourmate/payment-service/constant/env"
	grpc_connect "tourmate/payment-service/infrastructure/grpc"
	"tourmate/payment-service/infrastructure/grpc/user/pb"
	business_logic "tourmate/payment-service/interface/business_logic"

	"google.golang.org/grpc"
)

type userService struct {
	cnn    *grpc.ClientConn
	logger *log.Logger
}

func GenerateUserService(logger *log.Logger) (business_logic.IUserService, error) {
	cnn, err := grpc_connect.ConnectGrpcService(os.Getenv(env.USER_SERVICE_GRPC_PORT), os.Getenv(env.SERVICE_NAME), logger)

	if err != nil {
		return nil, err
	}

	return &userService{
		cnn:    cnn,
		logger: logger,
	}, nil
}

// GetUser implements businesslogic.IUserService.
func (u *userService) GetCustomerById(ctx context.Context, req *pb.GetCustomerByIdRequest) (*pb.CustomerResponse, error) {
	res, err := pb.NewUserServiceClient(u.cnn).GetCustomerById(ctx, req)

	if err != nil {
		u.logger.Println(err)
		return nil, err // Return the original gRPC error for the business logic to handle
	}

	return res, nil
}
