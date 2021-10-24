package jsonenc

import (
	"encoding/json"
	"net/http"
)

func JSONBalanceInfo(w http.ResponseWriter, id int64, userBalance float64) {
	type BalanceInformation struct {
		Id      int64   `json:"user id"`
		Balance float64 `json:"balance"`
	}

	balanceInfo := BalanceInformation{
		Id:      id,
		Balance: userBalance,
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(&balanceInfo)
	if err != nil {
		panic(err)
	}
}
