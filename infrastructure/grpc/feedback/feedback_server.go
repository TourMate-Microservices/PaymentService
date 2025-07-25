package feedback

import (
	"context"
	"tourmate/payment-service/infrastructure/grpc/feedback/pb"
	"tourmate/payment-service/repository"
	"tourmate/payment-service/repository/db"
	db_server "tourmate/payment-service/repository/db_server"
	"tourmate/payment-service/utils"
)

type feedbackGrpcService struct {
	pb.UnimplementedPaymentServiceServer
}

func GenerateGrpcService() *feedbackGrpcService {
	return &feedbackGrpcService{}
}

func (f *feedbackGrpcService) GetTourServiceRating(ctx context.Context, req *pb.GetTourServiceRatingRequest) (*pb.TourServiceRatingResponse, error) {
	var logger = utils.GetLogConfig()

	db, err := db.ConnectDB(logger, db_server.InitializePostgreSQL())
	if err != nil {
		return nil, err
	}

	rating, totalCount, err := repository.InitializeFeedbackRepo(db, logger).GetFeedbacksDetailByService(int(req.ServiceId), ctx)

	return &pb.TourServiceRatingResponse{
		Rating:      rating,
		ReviewCount: int32(totalCount),
	}, err
}
