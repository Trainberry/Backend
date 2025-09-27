package structures

import "github.com/go-ble/ble"

type Device struct {
	Name       string     `json:"name"`
	Connection ble.Client `json:"-"`
	ErrorCount int        `json:"error_count"`
	Speed      int        `json:"speed"`
	Light      bool       `json:"light"`
}

type Message struct {
	Operation string `json:"operation"`
	Device    Device `json:"device"`
}
