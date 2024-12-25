package structures

import (
	"time"
)

type Train struct {
	Name string `json:"name"`
	Model string `json:"model"`
	IP string `json:"ip"`
	LastContact time.Time `json:"last_contact"`
	LightsActivated bool `json:"lights_on"`
	Speed int `json:"speed"`
	IsGoingForward bool `json:"is_going_forward"`
}