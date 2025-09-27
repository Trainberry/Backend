package state

import (
	"github.com/rs/zerolog"
	"test/internal/config"
	"time"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(config.LogLevel)

	go func() {
		for {
			DevicesSanityCleanup()
			time.Sleep(config.SanityRunInterval)
		}
	}()

}
