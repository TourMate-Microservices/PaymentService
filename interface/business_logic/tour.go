package businesslogic

import (
	"context"
	"tourmate/payment-service/infrastructure/grpc/tour/pb"
)

type ITourService interface {
	GetTourById(ctx context.Context, req *pb.TourServiceIdRequest) (*pb.TourServiceItem, error)
}
