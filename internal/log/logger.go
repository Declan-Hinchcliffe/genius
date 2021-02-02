package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger = zap.NewNop().Sugar()

func InitLogger() error {
	core, err := newCore()
	if err != nil {
		return err
	}
	Logger = zap.New(core).Sugar()

	return nil
}

func newCore() (zapcore.Core, error) {
	getLevel := os.Getenv("LOG_LEVEL")
	var logLevel zapcore.Level

	err := logLevel.UnmarshalText([]byte(getLevel))
	if err != nil {
		return nil, err
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		os.Stdout,
		logLevel,
	)

	return core, nil
}
