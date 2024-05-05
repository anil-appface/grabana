package encoder

import (
	"github.com/anil-appface/grabana/encoder/golang"
	"github.com/anil-appface/sdk"
	"go.uber.org/zap"
)

func ToGolang(logger *zap.Logger, dashboard sdk.Board) (string, error) {
	golangEncoder := golang.NewEncoder(logger)

	return golangEncoder.EncodeDashboard(dashboard)
}
