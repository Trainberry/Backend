package controllers

import (
	"encoding/json"
	"net/http"
	"test/internal/state"
)

func GetState(w http.ResponseWriter, _ *http.Request) {
	cors(w)

	devices := state.Devices.GetDevices()
	d, _ := json.Marshal(devices)
	_, _ = w.Write(d)

}
