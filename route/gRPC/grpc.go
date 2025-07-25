package grpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"tourmate/payment-service/constant/env"
	"tourmate/payment-service/constant/noti"

	"google.golang.org/grpc"
)

func InitializeGRPCRoute(logger *log.Logger, service string) {
	var port string = os.Getenv(env.PAYMENT_SERVICE_GRPC_PORT)

	listen, err := net.Listen(os.Getenv(env.NETWORK), ":"+port)
	if err != nil {
		logger.Println(fmt.Sprintf(noti.NET_LISTENING_ERR_MSG, port) + err.Error())
		return
	}

	// rsServer, err := role_grpc.GenerateGRPCService()
	// if err != nil {
	// 	logger.Println(fmt.Sprintf(noti.GrpcGenerateMsg, service) + err.Error())
	// 	return
	// }

	var gRPCServer = grpc.NewServer()

	//pb.RegisterRoleServiceServer(gRPCServer, rsServer)

	log.Println(service + " gRPC starts listening on port " + port)
	if err := gRPCServer.Serve(listen); err != nil {
		logger.Println(fmt.Sprintf(noti.GRPC_CONNECTION_ERR_MSG, service) + err.Error())
		return
	}
}
