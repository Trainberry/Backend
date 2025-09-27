package controllers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"net/http"
	"test/internal/bluetooth"
	"test/internal/state"
	"test/internal/structures"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func Websocket(w http.ResponseWriter, r *http.Request) {
	cors(w)

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Err(err).Msg("upgrader.Upgrade failed")
		return
	}
	defer c.Close()

	go func() {
		for {
			if err := readMessage(c); err != nil {
				return
			}
		}
	}()

	for {
		if err := sendMessage(c); err != nil {
			return
		}
	}
}

func sendMessage(c *websocket.Conn) error {
	if err := c.WriteMessage(websocket.TextMessage, <-state.Messages); err != nil {
		log.Err(err).Msg("WriteMessage failed")
		return err
	}
	return nil
}

func readMessage(c *websocket.Conn) error {
	_, message, err := c.ReadMessage()
	if err != nil {
		log.Err(err).Msg("ReadMessage failed")
		return err
	}

	var parsed structures.Message
	_ = json.Unmarshal(message, &parsed)

	dev := state.Devices.GetDevice(parsed.Device.Name)
	if dev == nil {
		log.Err(err).Msg("unknown device")
		return nil
	}

	if parsed.Operation == "set_light" {
		_ = bluetooth.WriteLight(parsed.Device.Light, *dev)
	} else if parsed.Operation == "set_speed" {
		_ = bluetooth.WriteSpeed(parsed.Device.Speed, *dev)
	} else if parsed.Operation == "emerg_stop" {
		for _, k := range state.Devices.GetDevices() {
			go bluetooth.WriteSpeed(0, k)
		}

	} else {
		log.Err(err).Msg("unknown operation")
		return nil
	}

	log.Info().Msgf("Message received: %s", message)

	return nil
}
