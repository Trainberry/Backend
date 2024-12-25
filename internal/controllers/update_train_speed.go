package controllers

import(
	"net/http"
	"io"
	"github.com/sirupsen/logrus"
	"github.com/go-chi/chi/v5"
	"server/internal/services"
	"encoding/json"
)

func UpdateTrainSpeed(w http.ResponseWriter, r *http.Request) {
	trainName := chi.URLParam(r, "name")
	if trainName == "" {
		w.WriteHeader(404)
		return
	}

	// Read payload to collect data
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("Tried to read body but an error occurred: %v", err)
		w.WriteHeader(400)
		return
	}

	var data updatePayload
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		logrus.Errorf("Tried to unmarshald %s to updatePayload struct but an error occurred: %v", string(bodyBytes), err)
		w.WriteHeader(400)
		return
	}

	if data.Speed < 0 || data.Speed > 100 {
		logrus.Errorf("Invalid speed %d", data.Speed)
		w.WriteHeader(400)
		return
	}

	err = services.UpdateTrainSpeed(trainName, data.IsGoingForward, data.Speed)
	if err != nil {
		logrus.Errorf("Error while setting speed for train %s in service: %v", trainName, err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)

}

type updatePayload struct {
	Speed int `json:"speed"`
	IsGoingForward bool `json:"is_going_forward"`
}
