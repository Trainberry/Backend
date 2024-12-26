package controllers

import(
	"net/http"
	"io"
	"server/internal/structures"
	"server/internal/services"
	"encoding/json"
)

func RegisterTrain(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	var train structures.Train
	_ = json.Unmarshal(data, &train)

	if err = services.RegisterTrain(train); err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(201)

}