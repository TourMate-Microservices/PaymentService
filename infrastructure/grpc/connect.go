package grpc

import (
	"errors"
	"fmt"
	"log"
	"tourmate/payment-service/constant/noti"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectGrpcService(port, service string, logger *log.Logger) (*grpc.ClientConn, error) {
	cnn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		logger.Println(fmt.Sprintf(noti.GRPC_CONNECTION_ERR_MSG, service) + err.Error())
		return nil, errors.New(noti.INTERNALL_ERR_MSG)
	}

	return cnn, nil
}
