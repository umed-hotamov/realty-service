package logger

import (
	"log"
	"os"
	"path/filepath"

	"github.com/umed-hotamov/realty-service/internal/app/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func prepareLogPath(cfg *config.Config) error {
  _, err := os.Stat(cfg.Logger.Path)
  if os.IsNotExist(err) {
    err := os.MkdirAll(cfg.Logger.Path, 0777)
    if err != nil {
      return err
    }
  }

  return nil
}

func openLogFile(cfg *config.Config) (*os.File, error) {
  file, err := os.OpenFile(filepath.Join(cfg.Logger.Path, cfg.Logger.FileName), 
                           os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
  if err != nil {
    return nil, err
  }

  return file, nil
} 

func NewLogger(cfg *config.Config) *zap.Logger {
  err := prepareLogPath(cfg)
  if err != nil {
    log.Fatalf("failed to prepare log path: %v", err)
  }
  logFile, err := openLogFile(cfg)
  if err != nil {
    log.Fatalf("failed to open log file: %v", err)
  }
  
  var logLevel zapcore.Level
  switch cfg.Logger.Level {
  case "debug":
    logLevel = zapcore.DebugLevel
  case "warn": 
    logLevel = zapcore.WarnLevel
  case "info":
    logLevel = zapcore.InfoLevel
  case "error":
    logLevel = zapcore.ErrorLevel
  case "fatal":
    logLevel = zapcore.FatalLevel
  default:
    log.Fatalf("invalid log level provided: %s", cfg.Logger.Level)
  }
  
  encoderConfig := zap.NewProductionEncoderConfig()
  encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
  encoderConfig.TimeKey = "timestamp"
  
  encoder := zapcore.NewJSONEncoder(encoderConfig)
  writer := zapcore.AddSync(logFile)
  core := zapcore.NewTee(
    zapcore.NewCore(encoder, writer, logLevel),
  )

  return zap.New(core, zap.AddCaller())
}
