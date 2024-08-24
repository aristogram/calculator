package handler

import (
	"context"
	calcv1 "github.com/aristogram/protos/gen/go/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Calculator interface {
	Calculate(
		ctx context.Context,
		expr *calcv1.ExpressionRequest,
	) (string, error)
}

type serverAPI struct {
	calcv1.UnimplementedCalcServer
	Calculator
}

func Register(gRPC *grpc.Server) {
	calcv1.RegisterCalcServer(gRPC, serverAPI{})
}

func (s serverAPI) CalcExpr(
	ctx context.Context,
	expr *calcv1.ExpressionRequest,
) (*calcv1.Answer, error) {
	ans, err := s.Calculate(ctx, expr)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid expression")
	}

	return &calcv1.Answer{Answer: ans}, nil
}
