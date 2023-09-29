package main

import (
	"context"
	"dirScanner/internal/client/config"
	dirscanner "dirScanner/internal/server/proto"
	"dirScanner/internal/server/storage"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"sync"
)

func main() {
	clientConfig := config.NewConfig()
	ctx := context.Background()
	conn, err := grpc.Dial(
		clientConfig.GRPC,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	dirscannerClient := dirscanner.NewDirscannerClient(conn)
	var cacheMap = storage.NewCacheStorage()

	for {
		var path string
		fmt.Println("Enter directory path : ")
		_, err = fmt.Fscan(os.Stdin, &path)
		if err != nil {
			log.Fatal(err)
		}

		wg := new(sync.WaitGroup)

		_, ok := cacheMap.GetValue(path)
		if !ok {
			wg.Add(1)
			go func() {
				defer wg.Done()
				fmt.Println("cache doesn't have value, need to request")
				dirResponse, err := dirscannerClient.Scan(ctx, &dirscanner.DirScanRequest{
					Path: path,
				})
				if err != nil {
					log.Fatal(err)
				}
				cacheMap.SetValue(path, dirResponse.Files)
			}()
		}
		wg.Wait()
		cacheMap.Print(path)
	}
}
