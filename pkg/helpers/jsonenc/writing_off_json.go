package jsonenc

import (
	"encoding/json"
	"net/http"
)

func JSONWritingOff(w http.ResponseWriter, id int64, amount float64) {
	type BalanceInformation struct {
		Id     int64   `json:"user id"`
		Amount float64 `json:"writing off an amount"`
	}

	upBalanceInfo := BalanceInformation{
		Id:     id,
		Amount: amount,
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(&upBalanceInfo)
	if err != nil {
		panic(err)
	}
}
