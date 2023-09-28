package grpc

import (
	grpc "dirScanner/internal/server/proto"
)

type Handler struct {
	grpc.UnimplementedDirscannerServer
}

func NewHandler() *Handler {
	return &Handler{}
}
