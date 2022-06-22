package mylog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

func init() {
	infoLogFile := "./log/info.log"
	errorLogFile := "./log/error.log"

	_, err := os.Stat(infoLogFile)
	if os.IsNotExist(err) {
		_, err := os.Create(infoLogFile)
		if err != nil {
			panic(any(err))
		}
	}

	_, err = os.Stat(errorLogFile)
	if os.IsNotExist(err) {
		_, err := os.Create(errorLogFile)
		if err != nil {
			panic(any(err))
		}
	}

	logger, err := getLogger(infoLogFile, errorLogFile)
	if err != nil {
		panic(any(err))
	}
	Logger = logger
}

func getLogger(infoPath, errorPath string) (*zap.Logger, error) {
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	prodEncoder := zap.NewProductionEncoderConfig()
	prodEncoder.EncodeTime = zapcore.ISO8601TimeEncoder

	lowWriteSyncer, lowClose, err := zap.Open(infoPath)
	if err != nil {
		lowClose()
		return nil, err
	}

	highWriteSyncer, highClose, err := zap.Open(errorPath)
	if err != nil {
		highClose()
		return nil, err
	}

	highCore := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), highWriteSyncer, highPriority)
	lowCore := zapcore.NewCore(zapcore.NewJSONEncoder(prodEncoder), lowWriteSyncer, lowPriority)

	return zap.New(zapcore.NewTee(highCore, lowCore), zap.AddCaller()), nil
}
