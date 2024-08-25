package app

import (
	grpcapp "calculator/internal/app/grpc"
	"calculator/internal/services/calc"
	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, port int) *App {
	calcService := calc.New(log)

	grpcApp := grpcapp.New(log, calcService, port)

	return &App{
		GRPCServer: grpcApp,
	}
}
