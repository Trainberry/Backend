package bluetooth

import (
	"fmt"
	"github.com/go-ble/ble"
	"github.com/rs/zerolog/log"
	"strconv"
	"test/internal/state"
	"test/internal/structures"
	"time"
)

// ReadLightState returns false if light is off, on either.
func ReadLightState(device structures.Device) (bool, error) {
	if val, err := readInformation(LightInformation, device); err != nil {
		log.Error().Err(err).Msgf("Failed to read light info for device %s", device.Name)
		return false, err
	} else {
		light := false
		if string(val) == "on" {
			light = true
		}
		return light, nil
	}
}

// ReadSpeedState returns an int between -100 and 100. -100 is full-backward, 100 is full-forward.
func ReadSpeedState(device structures.Device) (int, error) {
	if val, err := readInformation(SpeedInformation, device); err != nil {
		log.Error().Err(err).Msgf("Failed to read speed info for device %s", device.Name)
		return 0, err
	} else {
		speed, _ := strconv.Atoi(string(val))
		return speed, nil
	}
}

func readInformation(desiredInformation int, device structures.Device) ([]byte, error) {
	if desiredInformation != PingInformation && desiredInformation != SpeedInformation && desiredInformation != LightInformation {
		log.Error().Msg("Desired information is not valid")
		return nil, fmt.Errorf("desired information is not valid")
	}

	data, err := device.Connection.ReadCharacteristic(&ble.Characteristic{ValueHandle: uint16(desiredInformation)})
	if err != nil {
		log.Error().Err(err).Msg("Failed to read desired information")
		state.Devices.IncrementDeviceFailCount(device.Name)
		return nil, err
	} else {
		state.Devices.ResetDeviceFailCount(device.Name)
	}

	time.Sleep(100 * time.Millisecond)

	return data, nil
}
