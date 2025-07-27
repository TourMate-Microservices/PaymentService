package tour

import (
	"context"
	"errors"
	"log"
	"os"
	"tourmate/payment-service/constant/env"
	"tourmate/payment-service/constant/noti"
	grpc_connect "tourmate/payment-service/infrastructure/grpc"
	"tourmate/payment-service/infrastructure/grpc/tour/pb"
	business_logic "tourmate/payment-service/interface/business_logic"

	"google.golang.org/grpc"
)

type tourService struct {
	cnn    *grpc.ClientConn
	logger *log.Logger
}

func GenerateTourService(logger *log.Logger) (business_logic.ITourService, error) {
	cnn, err := grpc_connect.ConnectGrpcService(os.Getenv(env.TOUR_SERVICE_GRPC_PORT), os.Getenv(env.SERVICE_NAME), logger)

	if err != nil {
		return nil, err
	}

	return &tourService{
		cnn:    cnn,
		logger: logger,
	}, nil
}

// GetTour implements businesslogic.ITourService.
func (t *tourService) GetTourById(ctx context.Context, req *pb.TourServiceIdRequest) (*pb.TourServiceItem, error) {
	res, err := pb.NewTourServiceClient(t.cnn).GetTourById(ctx, req)

	if err != nil {
		t.logger.Println(err)
		return nil, errors.New(noti.INTERNALL_ERR_MSG)
	}

	return res, nil
}
