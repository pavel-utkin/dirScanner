package grpc

import (
	grpc "dirScanner/internal/server/proto"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger *logrus.Logger
	grpc.UnimplementedDirscannerServer
}

// NewHandler - creates a new grpc server instance
func NewHandler(log *logrus.Logger) *Handler {
	return &Handler{logger: log}
}
