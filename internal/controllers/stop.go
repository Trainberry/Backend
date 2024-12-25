package controllers

import(
	"net/http"
	"server/internal/services"
)

func Stop(w http.ResponseWriter, r *http.Request) {
	services.StopAll()
}