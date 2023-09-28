package main

import (
	"dirScanner/internal/client/config"
	dirscanner "dirScanner/internal/server/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {

	clientConfig := config.NewConfig()

	conn, err := grpc.Dial(
		clientConfig.GRPC,
	)
	if err != nil {
		log.Fatal(err)
	}
	dirscannerClient := dirscanner.NewDirscannerClient(conn)
	fmt.Println(dirscannerClient)
	/*
		dirResponse, err := dirscannerClient.Scan(dirscanner.DirScanRequest{
			Path: ".",
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dirResponse)
	*/
}
