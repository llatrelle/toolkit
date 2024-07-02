package logger

import (
	"github.com/rs/zerolog/log"
)

func Info(msg string) {
	log.Info().Msg(msg)
}

func Debug(msg string) {
	log.Debug().Msg(msg)
}

func Error(msg string, err error) {
	log.Error().Err(err).Msg(msg)
}
