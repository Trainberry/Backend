package bluetooth

import (
	"errors"
	"github.com/go-ble/ble"
	"github.com/rs/zerolog/log"
	"strconv"
	"test/internal/state"
	"test/internal/structures"
	"time"
)

// WriteSpeed takes an integer between -100 and 100 and communicate to the chip
func WriteSpeed(speed int, device structures.Device) error {
	if speed < -100 || speed > 100 {
		return errors.New("speed out of range")
	}

	parsed := []byte(strconv.Itoa(speed))
	if err := writeInformation(SpeedInformation, parsed, device); err != nil {
		return err
	} else {
		state.Devices.StoreSpeedState(device.Name, speed)
		return nil
	}
}

func WritePing(device structures.Device) (bool, error) {
	if err := writeInformation(PingInformation, []byte(""), device); err != nil {
		log.Error().Err(err).Msgf("Failed to read ping info for device %s", device.Name)
		return false, err
	}
	return true, nil
}

// WriteLight turns on the light if the first param is true, turns off the light if the first param is false.
func WriteLight(light bool, device structures.Device) error {
	parsed := []byte("off")
	if light {
		parsed = []byte("on")
	}
	if err := writeInformation(LightInformation, parsed, device); err != nil {
		return err
	} else {
		state.Devices.StoreLightState(device.Name, light)
		return nil
	}
}

func writeInformation(desiredInformation int, data []byte, device structures.Device) error {
	if desiredInformation != PingInformation && desiredInformation != SpeedInformation && desiredInformation != LightInformation {
		log.Error().Msg("Desired information is not valid")
		return errors.New("desired information is not valid")
	}

	err := device.Connection.WriteCharacteristic(&ble.Characteristic{ValueHandle: uint16(desiredInformation)}, data, true)
	if err != nil {
		log.Error().Err(err).Msg("Failed to write desired information")
		state.Devices.IncrementDeviceFailCount(device.Name)
		return err
	} else {
		state.Devices.ResetDeviceFailCount(device.Name)
	}

	time.Sleep(100 * time.Millisecond)

	return nil
}
