package user

import (
	"context"
	"errors"
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

// GetUser implements businesslogic.IUserService.
func (u *userService) GetUser(req pb.GetUserRequest, ctx context.Context) (*pb.User, error) {
	res, err := pb.NewUserServiceClient(u.cnn).GetUser(ctx, &req)

	if err != nil {
		u.logger.Println(err)
		return nil, errors.New(noti.INTERNALL_ERR_MSG)
	}

	return res, nil
}
