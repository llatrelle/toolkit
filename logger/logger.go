package logger

import (
	"github.com/rs/zerolog/log"
)

func Info(msg string) {
	log.Info().Msg(msg)
}

func Debug(msg string, value interface{}) {
	log.Debug().Interface("%#v", value).Msg(msg)
}

func Trace(contextInfo string, msg string, value interface{}) {
	log.Trace().Str("ctx", contextInfo).Interface("%#v", value).Msg(msg)
}

func Error(msg string, err error) {
	log.Error().Err(err).Msg(msg)
}
