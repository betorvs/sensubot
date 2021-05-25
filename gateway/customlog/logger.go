package customlog

import (
	"github.com/betorvs/sensubot/appcontext"
	"github.com/betorvs/sensubot/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//Logger wrapper with two Logger methods
type Logger struct {
	//Logger only supports structured logging (less features more performance)
	Logger *zap.Logger
	//Sugar supports structured and printf-style APIs (less performance more features)
	Sugar *zap.SugaredLogger
}

// Debug uses fmt.Sprint to construct and log a message.
func (logger Logger) Debug(args ...interface{}) {
	logger.Sugar.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func (logger Logger) Info(args ...interface{}) {
	logger.Sugar.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (logger Logger) Warn(args ...interface{}) {
	logger.Sugar.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (logger Logger) Error(args ...interface{}) {
	logger.Sugar.Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (logger Logger) DPanic(args ...interface{}) {
	logger.Sugar.DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (logger Logger) Panic(args ...interface{}) {
	logger.Sugar.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (logger Logger) Fatal(args ...interface{}) {
	logger.Sugar.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (logger Logger) Debugf(template string, args ...interface{}) {
	logger.Sugar.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (logger Logger) Infof(template string, args ...interface{}) {
	logger.Sugar.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (logger Logger) Warnf(template string, args ...interface{}) {
	logger.Sugar.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (logger Logger) Errorf(template string, args ...interface{}) {
	logger.Sugar.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (logger Logger) DPanicf(template string, args ...interface{}) {
	logger.Sugar.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (logger Logger) Panicf(template string, args ...interface{}) {
	logger.Sugar.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (logger Logger) Fatalf(template string, args ...interface{}) {
	logger.Sugar.Fatalf(template, args...)
}

//Sync flushes the log if needed
func (logger Logger) Sync() {
	_ = logger.Sugar.Sync()
}

func discoverLogLevel() zapcore.Level {
	switch config.Values.LogLevel {
	case "DEBUG":
		return zap.DebugLevel
	case "INFO":
		return zap.InfoLevel
	case "WARNING":
		return zap.WarnLevel
	case "ERROR":
		return zap.ErrorLevel
	case "PANIC":
		return zap.PanicLevel
	case "FATAL":
		return zap.FatalLevel
	}
	return zap.InfoLevel
}

//initLogger lazily loads a Logger
func initLogger() appcontext.Component {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(discoverLogLevel()),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    encoderCfg,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, _ := config.Build()
	sugar := logger.Sugar()
	return Logger{Logger: logger, Sugar: sugar}
}

func init() {
	appcontext.Current.Add(appcontext.Logger, initLogger)
}
