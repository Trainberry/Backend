package bluetooth

import (
	"context"
	"github.com/go-ble/ble"
	"github.com/rs/zerolog/log"
	"slices"
	"strings"
	"test/internal/state"
	"test/internal/structures"
	"time"
)

// DetectDevices find all devices starting with "Trainberry::" and connect to them.
func DetectDevices() {
	log.Debug().Msg("Starting bluetooth detection")
	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), 500*time.Millisecond))

	devices, _ := ble.Find(ctx, false, func(a ble.Advertisement) bool {
		return strings.HasPrefix(a.LocalName(), "Trainberry::")
	})

	var filtered []ble.Advertisement

	for _, k := range devices {
		if !slices.Contains(filtered, k) {
			filtered = append(filtered, k)
		}
	}

	for _, k := range filtered {
		ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), 15*time.Hour))

		cln, err := ble.Connect(ctx, func(a ble.Advertisement) bool {
			return a.LocalName() == k.LocalName()
		})
		if err != nil {
			log.Error().Err(err).Msgf("Failed to connect to device %s", k.LocalName())
		} else {
			// Remove old peripherical if already connected
			state.Devices.DeleteDevice(k.LocalName())

			// Append new
			dev := structures.Device{
				Name:       k.LocalName(),
				Connection: cln,
			}

			// Save device
			state.Devices.AddDevice(dev)

			// Get and store state
			if val, err := ReadSpeedState(dev); err == nil {
				state.Devices.StoreSpeedState(dev.Name, val)
			}

			if val, err := ReadLightState(dev); err == nil {
				state.Devices.StoreLightState(dev.Name, val)
			}

		}
	}
	log.Debug().Msgf("Bluetooth detection done. %d new devices connected", len(filtered))
}
