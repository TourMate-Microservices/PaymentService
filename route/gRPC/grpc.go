package grpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"tourmate/payment-service/constant/env"
	"tourmate/payment-service/constant/noti"
	"tourmate/payment-service/infrastructure/grpc/feedback"
	"tourmate/payment-service/infrastructure/grpc/feedback/pb"

	"google.golang.org/grpc"
)

func InitializeGRPCRoute(logger *log.Logger, service string) {
	var port string = os.Getenv(env.PAYMENT_SERVICE_GRPC_PORT)

	listen, err := net.Listen(os.Getenv(env.NETWORK), ":"+port)
	if err != nil {
		logger.Println(fmt.Sprintf(noti.NET_LISTENING_ERR_MSG, port) + err.Error())
		return
	}

	var gRPCServer = grpc.NewServer()

	pb.RegisterPaymentServiceServer(gRPCServer, feedback.GenerateGrpcService())

	log.Println(service + " gRPC starts listening on port " + port)
	if err := gRPCServer.Serve(listen); err != nil {
		logger.Println(fmt.Sprintf(noti.GRPC_CONNECTION_ERR_MSG, service) + err.Error())
		return
	}
}
