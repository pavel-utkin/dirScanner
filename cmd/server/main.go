package main

import (
	grpchandler "dirScanner/internal/server/api/grpc"
	dirscanner "dirScanner/internal/server/proto"
	"fmt"
	"log"
)

func main() {
	handlerGrpc := grpchandler.NewHandler()

	dirResponse, err := handlerGrpc.Scan(dirscanner.DirScanRequest{
		Path: ".",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dirResponse)
}
