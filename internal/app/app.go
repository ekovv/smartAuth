package app

import (
	"log/slog"
	grpcapp "smartAuth/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int) *App {
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
