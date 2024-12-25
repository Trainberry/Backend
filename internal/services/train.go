package services

import (
	"server/internal/structures"
	"server/internal/cache"
	"errors"
	"time"
	"github.com/sirupsen/logrus"
	"net/http"
	"fmt"
	"strconv"
)


func RegisterTrain(train structures.Train) error {
	if train.Name == "" || train.Model == "" || train.IP == "" {
		return errors.New("invalid payload")
	}

	if _, ok := cache.Map.Get(train.Name); !ok {
		logrus.Infof("Registering new train named %s", train.Name)
	} else {
		logrus.Infof("Already known train %s registered again! Welcome back.", train.Name)
	}

	train.LastContact = time.Now()
	train.LightsActivated = true
	train.Speed = 0
	train.IsGoingForward = true

	cache.Map.Set(train.Name, train)

	return nil
}

func GetTrains() map[string]structures.Train {
	return cache.Map.Items()
}

func RemoveTrain(trainName string) {
	cache.Map.Remove(trainName)
}

func UpdateLights(trainName string, lightsOn bool) error {
	item, ok := cache.Map.Get(trainName)
	if !ok {
		logrus.Infof("Train %s does not exists, could not change lights", trainName)
		return errors.New("train does not exists")
	}

	meth := http.MethodDelete
	if lightsOn {
		meth = http.MethodPost
	}

	// Create request
	req, err := http.NewRequest(meth, fmt.Sprintf("http://%s/lights", item.IP), nil)
	if err != nil {
		logrus.Errorf("Could not create request to set lights for train %s : %v", trainName, err)
		return err
	}

	// Send request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Errorf("Unable to send request to set lights for train %s : %v", trainName, err)
		return err
	}

	if res.StatusCode != 204 {
		logrus.Errorf("Expected status code 204 on set lights for train %s but got %d", trainName, res.StatusCode)
		return errors.New("invalid response code")
	}

	// Update state
	item.LightsActivated = lightsOn
	item.LastContact = time.Now()
	cache.Map.Set(trainName, item)

	return nil
}

func UpdateTrainSpeed(trainName string, goingForward bool, speed int) error {
	item, ok := cache.Map.Get(trainName)
	if !ok {
		logrus.Infof("Train %s does not exists, could not change speed", trainName)
		return errors.New("train does not exists")
	}

	// Create request
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/speed", item.IP), nil)
	if err != nil {
		logrus.Errorf("Could not create request to set speed for train %s : %v", trainName, err)
		return err
	}

	// Set headers
	req.Header.Set("Speed", strconv.Itoa(speed))

	dir := "0"
	if goingForward {
		dir = "1"
	}

	req.Header.Set("Direction", dir)

	// Send request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logrus.Errorf("Unable to send request to set speed for train %s : %v", trainName, err)
		return err
	}

	if res.StatusCode != 204 {
		logrus.Errorf("Expected status code 204 on set speed for train %s but got %d", trainName, res.StatusCode)
		return errors.New("invalid response code")
	}

	// Update state
	item.Speed = speed
	item.IsGoingForward = goingForward
	item.LastContact = time.Now()
	cache.Map.Set(trainName, item)

	return nil
}

func StopAll() {
	for _, k := range cache.Map.Keys() {
		go UpdateTrainSpeed(k, true, 0)
	}
}