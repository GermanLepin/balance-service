package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"tech_task"
)

var (
	asc            = "asc"
	desc           = "desc"
	description    = "description"
	senderReceiver = "sender receiver"
	refill         = "refill"
	FALSE          = "F"
	TRUE           = "T"
	nilValue       = ""
	data           = "created at"
	amount         = "amount"
	sortBy         = "sort by"
	orderBy        = "order by"
	ctx            = context.Background()
	id             = "user id"
	id1            = "user id1"
	id2            = "user id2"
	currency       = "currency"
	RUB            = "RUB"
	USD            = "USD"
	static         = 100.00
	mapCur         map[string]json.RawMessage
	valCur         map[string]json.RawMessage
	rates          = "rates"
	accessKey      = "27c4039d0e33e2f74fbdc7afa63c08a8"
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

func IdValidate(idAccount string) (int64, error) {
	id, err := strconv.ParseInt(idAccount, 0, 64)
	if err != nil {
		logrus.WithError(err).Errorf("error with parcing user id")
		return 0, errors.New("error with parcing user id")
	}
	if id < 1 {
		logrus.Errorf("incorrect value user id ")
		return 0, errors.New("incorrect value user id")
	}

	return id, nil
}

func AmountValidate(amountS string) (float64, error) {
	validAmount := strings.Split(amountS, ".")
	if len(validAmount) > 1 {
		if len(validAmount[1]) > 2 {
			logrus.Errorf("The amount have more then 2 decimal places")
			return 0, errors.New("the amount have more then 2 decimal places")
		}
	}

	amount, err := strconv.ParseFloat(amountS, 64)
	if err != nil {
		logrus.WithError(err).Errorf("Error with parcing amount")
		return 0, errors.New("error with parcing amount")
	}

	if amount < 0.01 {
		logrus.Errorf("The amount is negative")
		return 0, errors.New("the amount is negative")
	}

	return amount, nil
}

func Pars(r *http.Request, value string) (correctVal string) {
	r.ParseForm()
	paramsRequest := r.Form
	valueSlice := paramsRequest[value]
	correctValue := strings.Join(valueSlice, " ")
	return correctValue
}

func ParsJSON(r *http.Request) (map[string]string, error) {
	var mapUser map[string]string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.WithError(err).Errorf("error parcing request")
		return nil, errors.New("error with parcing user id")
	}

	if err := json.Unmarshal(body, &mapUser); err != nil {
		logrus.WithError(err).Errorf("error parcing JSON")
		return nil, errors.New("error parcing JSON")
	}

	return mapUser, nil
}

func GetConvertValue(w http.ResponseWriter, currency string) float64 {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://api.exchangeratesapi.io/v1/latest?access_key=%s&symbols=%s", accessKey, currency), nil)
	if err != nil {
		logrus.Errorf("Creating request error")
		w.WriteHeader(http.StatusInternalServerError)
	}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Printf("Error while request to auth service")
		w.WriteHeader(http.StatusUnauthorized)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Printf("Error while read response body")
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.Unmarshal(body, &mapCur); err != nil {
		logrus.Printf("Error parcing JSON")
	}

	mapCurrency := mapCur[rates]

	if err := json.Unmarshal(mapCurrency, &valCur); err != nil {
		logrus.Printf("Error parcing JSON")
	}

	srtingValueCur := string(valCur[currency][:])

	valueCurrency, err := strconv.ParseFloat(srtingValueCur, 64)
	if err != nil {
		logrus.Printf("Error parcing Float")
	}
	return valueCurrency
}

func UsdAmount(usdToEur, rub, amount float64) float64 {
	usdToRub := rub / usdToEur
	usdAmount := amount / usdToRub
	return usdAmount
}

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

func JSONGetDescriptions(w http.ResponseWriter, row tech_task.Description) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&row)
	if err != nil {
		logrus.WithError(err).Errorf(err.Error())
		return err
	}

	return nil
}

func JSONBalanceInfo(w http.ResponseWriter, id int64, userBalance float64) error {
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
		logrus.WithError(err).Errorf(err.Error())
		return err
	}

	return nil
}

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

func JSONWritingOff(w http.ResponseWriter, id int64, amount float64) error {
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
		logrus.WithError(err).Errorf(err.Error())
		return err
	}

	return nil
}
