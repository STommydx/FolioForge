package logger

import (
	"fmt"
	"os"
	"path"

	"github.com/fatih/color"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LogLevel string

var (
	DebugLevel  = LogLevel(zapcore.DebugLevel.String())
	InfoLevel   = LogLevel(zapcore.InfoLevel.String())
	WarnLevel   = LogLevel(zapcore.WarnLevel.String())
	ErrorLevel  = LogLevel(zapcore.ErrorLevel.String())
	DPanicLevel = LogLevel(zapcore.DPanicLevel.String())
	PanicLevel  = LogLevel(zapcore.PanicLevel.String())
	FatalLevel  = LogLevel(zapcore.FatalLevel.String())
)

var logLevelMap = map[LogLevel]zapcore.Level{
	DebugLevel:  zapcore.DebugLevel,
	InfoLevel:   zapcore.InfoLevel,
	WarnLevel:   zapcore.WarnLevel,
	ErrorLevel:  zapcore.ErrorLevel,
	DPanicLevel: zapcore.DPanicLevel,
	PanicLevel:  zapcore.PanicLevel,
	FatalLevel:  zapcore.FatalLevel,
}

type Config struct {
	Path  string `env:"PATH, default=./logs"`
	Level struct {
		File    LogLevel `env:"FILE, default=info"`
		Console LogLevel `env:"CONSOLE, default=info"`
	} `env:", prefix=LEVEL_"`
}

var Module = fx.Options(
	fx.Provide(New),
	fx.Provide(NewGormLogger),
	fx.WithLogger(func(logger *zap.SugaredLogger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: logger.Desugar()}
	}),
)

func New(config *Config) (*zap.SugaredLogger, error) {
	// create directory if needed
	err := os.MkdirAll(config.Path, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("error in creating log file folder for writing: %w", err)
	}

	// create a new writer for log rotation
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename: path.Join(config.Path, "server.log"),
	})

	// setting the log level for file/console log output
	fileLogLevel := logLevelMap[config.Level.File]
	consoleLogLevel := logLevelMap[config.Level.Console]

	// setup the encoders
	fileEncoderConfig := zap.NewProductionEncoderConfig()
	fileEncoder := zapcore.NewJSONEncoder(fileEncoderConfig)
	consoleEncoderConfig := zap.NewProductionEncoderConfig()
	colorMap := map[zapcore.Level]*color.Color{
		zapcore.DebugLevel:  color.New(color.FgMagenta),
		zapcore.InfoLevel:   color.New(color.FgBlue),
		zapcore.WarnLevel:   color.New(color.FgYellow),
		zapcore.ErrorLevel:  color.New(color.FgRed),
		zapcore.DPanicLevel: color.New(color.FgRed, color.Bold),
		zapcore.FatalLevel:  color.New(color.FgRed, color.Bold),
		zapcore.PanicLevel:  color.New(color.FgRed, color.Bold),
	}
	consoleEncoderConfig.EncodeLevel = func(l zapcore.Level, pae zapcore.PrimitiveArrayEncoder) {
		// custom encoding of level string as [INFO] style
		pae.AppendString(colorMap[l].Sprintf("[%s]", l.CapitalString()))
	}
	consoleEncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	consoleEncoderConfig.EncodeCaller = func(ec zapcore.EntryCaller, pae zapcore.PrimitiveArrayEncoder) {
		// custom encoding of the caller, now is set to the trimmed file path
		pae.AppendString(ec.TrimmedPath())
	}
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	// create the two cores for the logger
	// when writing to a file, the *os.File need to be locked with Lock() for concurrent access
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, fileWriter, fileLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stderr), consoleLogLevel),
	)

	return zap.New(core, zap.AddCaller()).Sugar(), nil
}
