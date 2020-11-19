package tykdefs

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Hello() string {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Print("hello world")
	return "Hello, world."
}