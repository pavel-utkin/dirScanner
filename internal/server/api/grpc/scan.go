package grpc

import (
	"context"
	grpc "dirScanner/internal/server/proto"
	"os"
	"path/filepath"
)

func (h *Handler) Scan(ctx context.Context, request *grpc.DirScanRequest) (*grpc.DirScanResponse, error) {
	files := make([]*grpc.File, 10)
	currentDir := &grpc.DirScanResponse{Path: request.Path, Files: files}
	err := filepath.Walk(request.Path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			currentDir.Files = append(currentDir.Files, &grpc.File{FileName: path, Size: info.Size()})
			return nil
		})
	if err != nil {
		return &grpc.DirScanResponse{}, nil
	}
	return currentDir, nil
}
