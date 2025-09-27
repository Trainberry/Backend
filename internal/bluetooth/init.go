package bluetooth

import (
	"github.com/go-ble/ble"
	"github.com/go-ble/ble/linux"
	"github.com/rs/zerolog/log"
	"test/internal/config"
	"time"
)

func init() {
	d, err := linux.NewDevice()

	if err != nil {
		log.Fatal().Err(err).Msg("Adapter init failed")
	}
	ble.SetDefaultDevice(d)

	go func() {
		for {
			time.Sleep(config.BLEDetectionRunInterval)
			DetectDevices()
		}
	}()
}
