package json_responce

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func JSONU2U(w http.ResponseWriter, id1, id2 int64, amount float64) error {
	type BalanceInformation struct {
		Id1    int64   `json:"user id sender"`
		Amount float64 `json:"writing off an amount"`
		Id2    int64   `json:"user id recipient"`
	}

	upBalanceInfo := BalanceInformation{
		Id1:    id1,
		Amount: amount,
		Id2:    id2,
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(&upBalanceInfo)
	if err != nil {
		logrus.WithError(err).Errorf(err.Error())
		return err
	}

	return nil
}
