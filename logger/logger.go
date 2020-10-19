package logger

import (
"go.uber.org/zap"
"go.uber.org/zap/zapcore"
)

var(
	log *zap.Logger

)

func init(){

	var err error
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:"json",
		EncoderConfig:zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "time",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}


	if log, err = logConfig.Build(); err != nil{
		panic(err)
	}

}

func Info(msg string, tags...zap.Field){
	log.Info(msg, tags...)
	_ = log.Sync()
}

func Error(msg string, err error, tags...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(msg, tags...)
	_ = log.Sync()
}