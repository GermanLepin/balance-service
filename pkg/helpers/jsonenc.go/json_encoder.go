package jsonenc

import (
	"encoding/json"
	"net/http"
)

func JSONError(w http.ResponseWriter, error_str string) {
	type JSONErr struct {
		Error string `json:"error"`
	}

	error_json := JSONErr{
		Error: error_str,
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(&error_json)
	if err != nil {
		panic(err)
	}
}

func JSONDeleteUser(w http.ResponseWriter, id int64, status string) {
	type JSONErr struct {
		Id     int64  `json:"id"`
		Status string `json:"status"`
	}

	error_json := JSONErr{
		Id:     id,
		Status: status,
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(&error_json)
	if err != nil {
		panic(err)
	}
}

func JSONEncoder(w http.ResponseWriter, id int64, userBalance float64) {
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

func JSONUpBalance(w http.ResponseWriter, id int64, amount float64) {
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
		panic(err)
	}
}

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

func JSONU2U(w http.ResponseWriter, id1, id2 int64, amount float64) {
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
		panic(err)
	}
}
