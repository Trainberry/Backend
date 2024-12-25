package controllers

import(
	"net/http"
	"github.com/go-chi/chi/v5"
	"server/internal/services"
)

func DisableTrainLights(w http.ResponseWriter, r *http.Request) {
	trainName := chi.URLParam(r, "name")
	if trainName == "" {
		w.WriteHeader(404)
		return
	}

	if err := services.UpdateLights(trainName, false); err != nil {
		w.WriteHeader(500)
		return
	}


	w.WriteHeader(204)
}