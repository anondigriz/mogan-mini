package logger

import (
	"os"
	"path"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Zap *zap.Logger
}

func New() *Logger {
	lg := &Logger{}
	return lg
}

func (lg *Logger) Init(projectsPath string, debug bool) error {
	logsPath := path.Join(projectsPath, "log")
	err := createDir(logsPath)
	if err != nil {
		return err
	}

	logFileName := path.Join(logsPath, time.Now().Format("2006-01-02")+".txt")

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	writer := zapcore.AddSync(logFile)

	defaultLogLevel := zapcore.ErrorLevel
	if debug {
		defaultLogLevel = zapcore.DebugLevel
	}

	var tee []zapcore.Core
	tee = append(tee, zapcore.NewCore(fileEncoder, writer, defaultLogLevel))

	if debug {
		consoleEncoder := zapcore.NewConsoleEncoder(config)
		tee = append(tee, zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel))
	}

	core := zapcore.NewTee(tee...)
	lg.Zap = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return nil
}

func createDir(logPath string) error {
	err := os.MkdirAll(logPath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
