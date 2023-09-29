package main

import (
	"context"
	"dirScanner/internal/client/config"
	dirscanner "dirScanner/internal/server/proto"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func main() {

	clientConfig := config.NewConfig()

	conn, err := grpc.Dial(
		clientConfig.GRPC,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	dirscannerClient := dirscanner.NewDirscannerClient(conn)
	fmt.Println(dirscannerClient)

	ctx := context.Background()

	var cacheMap = make(map[string][]*dirscanner.File)

	for {
		var path string
		fmt.Fscan(os.Stdin, &path)
		dirResponse, err := dirscannerClient.Scan(ctx, &dirscanner.DirScanRequest{
			Path: path,
		})

		cacheMap[path] = dirResponse.Files
		println(cacheMap)
		println(json.Marshal(dirResponse.Files))

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(dirResponse)
	}
}
