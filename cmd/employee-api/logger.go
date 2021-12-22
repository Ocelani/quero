package main

import (
	"os"
	"quero2pay/pkg/employee"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// _log is a private instance of this API logging instance.
var _log zerolog.Logger = NewLogger()

// NewLogger instantiates and starts the logger of this API.
func NewLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	return log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}

func LogEmployee(emp *employee.Employee, info string) {
	_log.Info().
		Int("id", emp.ID).
		Str("name", emp.Name).
		Float64("sallary", emp.Sallary).
		Str("role", string(emp.Role)).
		Msg(info)
}
