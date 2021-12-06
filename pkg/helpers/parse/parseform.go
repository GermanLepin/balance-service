package parse

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	mapUser map[string]string
)

func Pars(r *http.Request, value string) (correctVal string) {
	r.ParseForm()
	paramsRequest := r.Form
	valueSlice := paramsRequest[value]
	correctValue := strings.Join(valueSlice, " ")
	return correctValue
}

func ParsJSON(r *http.Request) (map[string]string, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.WithError(err).Errorf("Error parcing request")
		return nil, errors.New("error with parcing id")
	}

	if err := json.Unmarshal(body, &mapUser); err != nil {
		logrus.WithError(err).Errorf("Error parcing JSON")
		return nil, errors.New("error parcing JSON")
	}

	return mapUser, nil
}
