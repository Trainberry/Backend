package controllers

import(
	"net/http"
	"github.com/go-chi/chi/v5"
	"server/internal/services"
)

func RemoveTrain(w http.ResponseWriter, r *http.Request) {
	trainName := chi.URLParam(r, "name")
	if trainName == "" {
		w.WriteHeader(404)
		return
	}

	services.RemoveTrain(trainName)

	w.WriteHeader(204)


}