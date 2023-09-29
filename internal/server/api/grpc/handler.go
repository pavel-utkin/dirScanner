package grpc

import (
	grpc "dirScanner/internal/server/proto"
	"dirScanner/internal/server/storage"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger  *logrus.Logger
	storage *storage.Storage
	grpc.UnimplementedDirscannerServer
}

// NewHandler - creates a new grpc server instance
func NewHandler(log *logrus.Logger) *Handler {
	return &Handler{logger: log}
}
