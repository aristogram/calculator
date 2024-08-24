package calc

import (
	"context"
	"log/slog"
)

type Calculator struct {
	log *slog.Logger
}

func New(log *slog.Logger) *Calculator {
	return &Calculator{
		log: log,
	}
}

func (c *Calculator) Calculate(
	ctx context.Context,
	expr string,
) (string, error) {
	panic("not implemented")
}
