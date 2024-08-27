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
	calc Calculator
}

func Register(gRPC *grpc.Server, calc Calculator) {
	calcv1.RegisterCalcServer(gRPC, serverAPI{calc: calc})
}

func (s serverAPI) CalcExpr(
	ctx context.Context,
	expr *calcv1.ExpressionRequest,
) (*calcv1.Answer, error) {
	ans, err := s.calc.Calculate(ctx, expr)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid expression")
	}

	return &calcv1.Answer{Answer: ans}, nil
}
