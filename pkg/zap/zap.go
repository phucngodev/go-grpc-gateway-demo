package zap

import (
	"fmt"
	"kratos-blog/internal/conf"
	"os"
	"path"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZapLogger struct {
	log  *zap.Logger
	Sync func() error
}

func NewLogger(c *conf.Log) *ZapLogger {
	encodeConfig := zapcore.EncoderConfig{
		LevelKey:     "level",
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encodeConfig)

	writerSyncer := &lumberjack.Logger{
		Filename:   path.Join(c.Path, "log.log"),
		MaxBackups: int(c.MaxBackups),
		MaxAge:     int(c.MaxAge),
		Compress:   true,
	}

	var cores []zapcore.Core
	fileCore := zapcore.NewCore(encoder, zapcore.AddSync(writerSyncer), zapcore.InfoLevel)
	cores = append(cores, fileCore)
	if c.Debug {
		consoleCore := zapcore.NewCore(encoder, zapcore.AddSync(zapcore.Lock(os.Stdout)), zapcore.InfoLevel)
		cores = append(cores, consoleCore)
	}

	core := zapcore.NewTee(cores...)
	logger := zap.New(core)
	return &ZapLogger{
		log:  logger,
		Sync: logger.Sync,
	}
}

func (l *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}
	switch level {
	case log.LevelDebug:
		l.log.Debug("", data...)
	case log.LevelInfo:
		l.log.Info("", data...)
	case log.LevelWarn:

		l.log.Warn("", data...)
	case log.LevelError:
		l.log.Error("", data...)
	}
	return nil
}
