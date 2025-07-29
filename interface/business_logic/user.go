package businesslogic

import (
	"context"
	"tourmate/payment-service/infrastructure/grpc/user/pb"
)

type IUserService interface {
	GetCustomerById(ctx context.Context, req *pb.GetCustomerByIdRequest) (*pb.CustomerResponse, error)
	GetTourGuideById(ctx context.Context, req *pb.GetTourGuideByIdRequest) (*pb.TourGuideResponse, error)
}
