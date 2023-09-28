package api

import (
	grpcHandler "dirScanner/internal/server/api/grpc"
	"dirScanner/internal/server/config"
	grpcDirscanner "dirScanner/internal/server/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

// StartGRPCService - starts the GRPC gophkeeper server
func StartGRPCService(grpcHandler *grpcHandler.Handler, config *config.Config) {
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", config.AddressGRPC)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcDirscanner.RegisterDirscannerServer(grpcServer, grpcHandler)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed gprc server: %v", err)
	}
}
