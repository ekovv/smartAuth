package app

import (
	"log/slog"
	grpcapp "smartAuth/internal/app/grpc"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	//	gRPCServer := grpc.NewServer()
	//
	//	authgrpc.Register(gRPCServer)
	//	return &App{
	//		log:        log,
	//		gRPCServer: gRPCServer,
	//		port:       port,
	//	}
	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
