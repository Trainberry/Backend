package config

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var LogLevel = zerolog.InfoLevel
var SanityRunInterval = time.Second * 5
var BLEDetectionRunInterval = time.Second * 2

func init() {

	if os.Getenv("TRAINBERRY_LOG_LEVEL") != "" {
		level, err := zerolog.ParseLevel(os.Getenv("TRAINBERRY_LOG_LEVEL"))
		if err == nil {
			LogLevel = level
		}
	}

	if os.Getenv("TRAINBERRY_SANITY_RUN_INTERVAL") != "" {
		d, err := time.ParseDuration(os.Getenv("TRAINBERRY_SANITY_RUN_INTERVAL"))
		if err == nil {
			SanityRunInterval = d
		}
	}

	if os.Getenv("TRAINBERRY_BLE_DETECTION_RUN_INTERVAL") != "" {
		d, err := time.ParseDuration(os.Getenv("TRAINBERRY_BLE_DETECTION_RUN_INTERVAL"))
		if err == nil {
			BLEDetectionRunInterval = d
		}
	}

}
