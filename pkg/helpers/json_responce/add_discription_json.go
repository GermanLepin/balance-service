package json_responce

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func JSONUAddDescription(w http.ResponseWriter, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) error {
	type DescriptionInformation struct {
		Id              int64   `json:"user id"`
		BalanceAtMoment float64 `json:"balance at moment"`
		CorectAmount    float64 `json:"amount"`
		Description     string  `json:"description of transaction"`
		SenderReceiver  string  `json:"sender or receiver"`
		Refill          string  `json:"refill the balance"`
	}

	upBalanceInfo := DescriptionInformation{
		Id:              id,
		BalanceAtMoment: balanceAtMoment,
		CorectAmount:    corectAmount,
		Description:     description,
		SenderReceiver:  senderReceiver,
		Refill:          refill,
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(&upBalanceInfo)
	if err != nil {
		logrus.WithError(err).Errorf(err.Error())
		return err
	}

	return nil
}
