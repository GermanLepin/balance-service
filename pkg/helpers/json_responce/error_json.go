package json_responce

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func JSONError(w http.ResponseWriter, errorStr string) error {
	type JSONErr struct {
		Error string `json:"error"`
	}

	errorJson := JSONErr{
		Error: errorStr,
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(&errorJson)
	if err != nil {
		logrus.WithError(err).Errorf(err.Error())
		return err
	}

	return nil
}
