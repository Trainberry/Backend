package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"time"
	"github.com/sirupsen/logrus"
	"server/internal/controllers/middlewares"
	"server/internal/controllers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(15 * time.Second))
	r.Use(middlewares.CORSManager)

	r.Post("/register", controllers.RegisterTrain)
	r.Get("/stop", controllers.Stop)
	r.Get("/trains", controllers.GetTrains)
	r.Delete("/trains/{name}", controllers.RemoveTrain)
	r.Put("/trains/{name}/speed", controllers.UpdateTrainSpeed)
	r.Post("/trains/{name}/lights", controllers.EnableTrainLights)
	r.Delete("/trains/{name}/lights", controllers.DisableTrainLights)

	logrus.Info("Listing on port 8080")
	
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		logrus.Fatalf("HTTP server hangup! Error was: %v", err)
	}
}