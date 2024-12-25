package controllers

import(
	"net/http"
	"server/internal/structures"
	"server/internal/services"
	"encoding/json"
)

func GetTrains(w http.ResponseWriter, r *http.Request) {

	// Transform map to simple array

	tmpMap := services.GetTrains()

	var tmpArray []structures.Train

	for _, k := range tmpMap {
		tmpArray = append(tmpArray, k)
	}

	if tmpArray == nil {
		tmpArray = []structures.Train{}
	}

	data, err := json.Marshal(tmpArray)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	_, _ = w.Write(data)

}