package businesslogic

import (
	"context"
	"tourmate/payment-service/infrastructure/grpc/user/pb"
)

type IUserService interface {
	GetUser(req pb.GetUserRequest, ctx context.Context) (*pb.User, error)
}
