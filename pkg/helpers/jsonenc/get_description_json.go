package jsonenc

import (
	"encoding/json"
	"net/http"
	"tech_task/internal/entities"
)

func JSONUGetUserIdDescription(w http.ResponseWriter, row entities.Description) {

	encoder := json.NewEncoder(w)
	err := encoder.Encode(&row)
	if err != nil {
		panic(err)
	}
}

func JSONUGetAllUsers(w http.ResponseWriter, row entities.Description) {

	encoder := json.NewEncoder(w)
	err := encoder.Encode(&row)
	if err != nil {
		panic(err)
	}
}
