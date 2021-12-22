package main

import (
	"os"
	"quero2pay/pkg/address"
	"quero2pay/pkg/business"

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

func LogBusiness(bsns *business.Business, info string) {
	_log.Info().
		Int("id", bsns.ID).
		Str("name", bsns.Name).
		Str("phone", bsns.Phone).
		Msg(info)
}

func LogAddress(addr *address.Address, info string) {
	_log.Info().
		Int("id", addr.ID).
		Str("zip_code", addr.ZipCode).
		Str("city", addr.City).
		Str("state", addr.State).
		Str("district", addr.District).
		Str("street", addr.Street).
		Int("number", addr.Number).
		Str("ddd", addr.DDD).
		Msg(info)
}
