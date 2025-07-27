package businesslogic

import (
	"context"
	"tourmate/payment-service/infrastructure/grpc/tour/pb"
)

type ITourService interface {
	GetTour(ctx context.Context, req *pb.TourServiceIdRequest) (*pb.TourServiceItem, error)
}
