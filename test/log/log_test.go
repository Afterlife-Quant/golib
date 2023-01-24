package log

import (
	"context"
	"testing"

	"github.com/Afterlife-Quant/golib/pkg/log"
	"go.uber.org/zap/zapcore"
)

func TestLogger(t *testing.T) {
	log.StartWithLevel("tmp/temp.log", zapcore.InfoLevel)
	log.Info(context.TODO(), "testlog")
}
