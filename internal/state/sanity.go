package state

import (
	"github.com/rs/zerolog/log"
)

func DevicesSanityCleanup() {
	log.Debug().Msgf("Sanity cleanup is running. %d items before cleanup", len(Devices.GetDevices()))
	for _, k := range Devices.GetDevices() {
		if k.ErrorCount >= 3 {
			Devices.DeleteDevice(k.Name)
		}
	}

	log.Debug().Msgf("Sanity cleanup is finished. %d items after cleanup", len(Devices.GetDevices()))
}
