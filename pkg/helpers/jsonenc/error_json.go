package jsonenc

import (
	"encoding/json"
	"net/http"
)

func JSONError(w http.ResponseWriter, errorStr string) {
	type JSONErr struct {
		Error string `json:"error"`
	}

	errorJson := JSONErr{
		Error: errorStr,
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(&errorJson)
	if err != nil {
		panic(err)
	}
}
