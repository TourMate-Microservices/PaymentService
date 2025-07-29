package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"tourmate/payment-service/constant/env"
	"tourmate/payment-service/constant/noti"
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

const service string = "user"

// GetUser implements businesslogic.IUserService.
func (u *userService) GetCustomerById(ctx context.Context, req *pb.GetCustomerByIdRequest) (*pb.CustomerResponse, error) {
	res, err := pb.NewUserServiceClient(u.cnn).GetCustomerById(ctx, req)

	if err != nil {
		u.logger.Println(fmt.Sprintf(noti.GRPC_CONNECTION_ERR_MSG, service) + err.Error())
		return nil, errors.New(noti.INTERNALL_ERR_MSG)
	}

	return res, nil
}

// GetTourGuideById implements businesslogic.IUserService.
func (u *userService) GetTourGuideById(ctx context.Context, req *pb.GetTourGuideByIdRequest) (*pb.TourGuideResponse, error) {
	res, err := pb.NewUserServiceClient(u.cnn).GetTourGuideById(ctx, req)

	if err != nil {
		u.logger.Println(fmt.Sprintf(noti.GRPC_CONNECTION_ERR_MSG, service) + err.Error())
		return nil, errors.New(noti.INTERNALL_ERR_MSG)
	}

	return res, nil
}
