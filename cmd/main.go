package main

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"test/internal/bluetooth"
	"test/internal/controllers"
	"test/internal/state"
	"time"
)

func main() {
	// Detect devices (also start init loop)
	bluetooth.DetectDevices()

	// Check that connect chips are still alive
	go startHealthcheck()

	// Create WebSocket server
	http.HandleFunc("/ws", controllers.Websocket)
	http.HandleFunc("/state", controllers.GetState)
	log.Info().Msg("Listening on websocket port 8080")
	log.Fatal().Err(http.ListenAndServe(":8080", nil))
}

func startHealthcheck() {
	for {
		time.Sleep(3 * time.Second)
		for _, k := range state.Devices.GetDevices() {
			_, _ = bluetooth.WritePing(k)
		}
	}
}
