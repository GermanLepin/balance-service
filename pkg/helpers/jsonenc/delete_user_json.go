package jsonenc

import (
	"encoding/json"
	"net/http"
)

func JSONDeleteUser(w http.ResponseWriter, id int64, status string) {
	type DeleteUser struct {
		Id     int64  `json:"id"`
		Status string `json:"status"`
	}

	deleteUser := DeleteUser{
		Id:     id,
		Status: status,
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(&deleteUser)
	if err != nil {
		panic(err)
	}
}
