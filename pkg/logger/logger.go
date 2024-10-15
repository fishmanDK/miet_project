package logger

import (
	"log/slog"
	"os"
	"time"
)

type Config struct {
	DevMode bool   `yaml:"devMode"`
	Encoder string `yaml:"encoder"`
}

type Logger interface {
	InitLogger()
	Debug(template string)
	Debugf(template string, args ...interface{})
	Info(template string)
	Infof(template string, args ...interface{})
	Warn(template string)
	Warnf(template string, args ...interface{})
	Error(template string)
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(template string)
	WithName(name string)
	HttpMiddlewareAccessLogger(method string, uri string, status int, size int64, time time.Duration)
	GrpcMiddlewareAccessLogger(method string, time time.Duration, metaData map[string][]string, err error)
	GrpcClientInterceptorLogger(method string, req interface{}, reply interface{}, time time.Duration, metaData map[string][]string, err error)
	KafkaProcessMessage(topic string, partition int, message string, workerID int, offset int64, time time.Time)
	KafkaLogCommittedMessage(topic string, partition int, offset int64)
}

type appLogger struct {
	devMode  bool
	encoding string
	// sugarLogger *zap.SugaredLogger //FIXME
	logger *slog.Logger
}

func NewAppLogger(cfg Config) *appLogger {
	return &appLogger{devMode: cfg.DevMode, encoding: cfg.Encoder}
}

func (l *appLogger) InitLogger() {
	var opts slog.HandlerOptions
	if l.devMode {
		opts = slog.HandlerOptions{Level: slog.LevelDebug}
	} else {
		opts = slog.HandlerOptions{Level: slog.LevelInfo}
	}

	var slogHandler slog.Handler
	if l.encoding == "console" {
		slogHandler = slog.NewTextHandler(os.Stdout, &opts)
	} else {
		slogHandler = slog.NewJSONHandler(os.Stdout, &opts)
	}

	l.logger = slog.New(slogHandler)
}

func (l *appLogger) Debug(template string) {
	l.logger.Debug(template)
}

func (l *appLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debug(template, args...)
}

func (l *appLogger) Info(template string) {
	l.logger.Info(template)
}

func (l *appLogger) Infof(template string, args ...interface{}) {
	l.logger.Info(template, args...)
}

func (l *appLogger) Warn(template string) {
	l.logger.Warn(template)
}

func (l *appLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warn(template, args...)
}

func (l *appLogger) Error(template string) {
	l.logger.Error(template)
}

func (l *appLogger) Errorf(template string, args ...interface{}) {
	l.logger.Error(template, args...)
}

func (l *appLogger) DPanic(args ...interface{}) {}

func (l *appLogger) DPanicf(template string, args ...interface{}) {}

func (l *appLogger) Fatal(template string) {
	l.Error(template)
	os.Exit(1)
}

func (l *appLogger) WithName(name string) {}

func (l *appLogger) HttpMiddlewareAccessLogger(method string, uri string, status int, size int64, time time.Duration) {
}

func (l *appLogger) GrpcMiddlewareAccessLogger(method string, time time.Duration, metaData map[string][]string, err error) {
}

func (l *appLogger) GrpcClientInterceptorLogger(method string, req interface{}, reply interface{}, time time.Duration, metaData map[string][]string, err error) {
}

func (l *appLogger) KafkaProcessMessage(topic string, partition int, message string, workerID int, offset int64, time time.Time) {
}

func (l *appLogger) KafkaLogCommittedMessage(topic string, partition int, offset int64) {}
