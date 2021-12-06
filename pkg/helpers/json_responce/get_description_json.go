package json_responce

import (
	"encoding/json"
	"net/http"
	"tech_task"

	"github.com/sirupsen/logrus"
)

func JSONUGetUserIdDescription(w http.ResponseWriter, row tech_task.Description) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&row)
	if err != nil {
		logrus.WithError(err).Errorf(err.Error())
		return err
	}

	return nil
}

func JSONUGetAllUsers(w http.ResponseWriter, row tech_task.Description) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&row)
	if err != nil {
		logrus.WithError(err).Errorf(err.Error())
		return err
	}

	return nil
}
