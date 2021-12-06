package json_responce

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func JSONUpBalance(w http.ResponseWriter, id int64, amount float64) error {
	type BalanceInformation struct {
		Id     int64   `json:"user id"`
		Amount float64 `json:"top up an amount"`
	}

	upBalanceInfo := BalanceInformation{
		Id:     id,
		Amount: amount,
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(&upBalanceInfo)
	if err != nil {
		logrus.WithError(err).Errorf(err.Error())
		return err
	}

	return nil
}
