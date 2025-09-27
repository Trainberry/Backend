package state

import (
	"encoding/json"
	"slices"
	"strings"
	"test/internal/structures"
)

var Devices = Device{
	devices: &[]structures.Device{},
}

// Device is a store that contains an out-of-range array of connected devices. Use methods below to interact with it.
type Device struct {
	devices *[]structures.Device
}

// AddDevice adds a new device
func (d Device) AddDevice(device structures.Device) {
	dev := append(*d.devices, device)
	*d.devices = dev
	d.notify("added_device", device)
}

// GetDevices returns a list of all devices connected to the server
func (d Device) GetDevices() []structures.Device {
	return *d.devices
}

// GetDevice return a device by its unique name (nil if not found)
func (d Device) GetDevice(name string) *structures.Device {
	for _, device := range *d.devices {
		if strings.HasSuffix(device.Name, name) {
			return &device
		}
	}
	return nil
}

// DeleteDevice delete a device by its name (if the device does not exist, it will not throw)
func (d Device) DeleteDevice(name string) {
	dev := slices.DeleteFunc(*d.devices, func(device structures.Device) bool {
		return device.Name == name
	})
	*d.devices = dev
	d.notify("deleted_device", structures.Device{Name: name})
}

// IncrementDeviceFailCount adds one to the failed count for this device, allowing the sanity check to delete them later
func (d Device) IncrementDeviceFailCount(name string) {
	dev := *d.devices
	for i, _ := range dev {
		if dev[i].Name == name {
			dev[i].ErrorCount += 1
			d.notify("fail_counter_increment", dev[i])
		}
	}
	*d.devices = dev
}

// ResetDeviceFailCount reset the failed count for this device (eg. the device is back online)
func (d Device) ResetDeviceFailCount(name string) {
	dev := *d.devices
	for i, _ := range dev {
		if dev[i].Name == name {
			if dev[i].ErrorCount > 0 {
				d.notify("fail_counter_reset", dev[i])
			}
			dev[i].ErrorCount = 0
		}
	}
	*d.devices = dev
}

// StoreLightState saves the current light state for the device
func (d Device) StoreLightState(name string, up bool) {
	dev := *d.devices
	for i, _ := range dev {
		if dev[i].Name == name {
			dev[i].Light = up
			d.notify("light_update", dev[i])
		}
	}
	*d.devices = dev
}

// StoreSpeedState saves the current speed state for the device
func (d Device) StoreSpeedState(name string, speed int) {
	dev := *d.devices
	for i, _ := range dev {
		if dev[i].Name == name {
			dev[i].Speed = speed
			d.notify("speed_update", dev[i])
		}
	}
	*d.devices = dev
}

func (d Device) notify(action string, device structures.Device) {
	go func() {
		msg := structures.Message{
			Operation: action,
			Device:    device,
		}
		payload, _ := json.Marshal(msg)
		Messages <- payload
	}()
}
