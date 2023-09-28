package grpc

import (
	"context"
	grpc "dirScanner/internal/server/proto"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func (h *Handler) Scan(ctx context.Context, request *grpc.DirScanRequest) (*grpc.DirScanResponse, error) {
	// scan dirs
	println(ctx)
	err := filepath.Walk(request.Path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size(), "Bytes")
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	return &grpc.DirScanResponse{}, nil
}
