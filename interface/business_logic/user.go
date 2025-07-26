package businesslogic

import (
	"context"
	"tourmate/payment-service/infrastructure/grpc/user/pb"
)

type IUserService interface {
	GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error)
}
