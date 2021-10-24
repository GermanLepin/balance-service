package convert

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var (
	mapCur    map[string]json.RawMessage
	valCur    map[string]json.RawMessage
	rates     = "rates"
	accessKey = "27c4039d0e33e2f74fbdc7afa63c08a8"
)

func GetConvertValue(w http.ResponseWriter, currency string) float64 {
	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://api.exchangeratesapi.io/v1/latest?access_key=%s&symbols=%s", accessKey, currency), nil)
	if err != nil {
		log.Errorf("Creating request error")
		w.WriteHeader(http.StatusInternalServerError)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error while request to auth service")
		w.WriteHeader(http.StatusUnauthorized)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error while read response body")
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.Unmarshal(body, &mapCur); err != nil {
		log.Printf("Error parcing JSON")
	}

	mapCurrency := mapCur[rates]

	if err := json.Unmarshal(mapCurrency, &valCur); err != nil {
		log.Printf("Error parcing JSON")
	}

	srtingValueCur := string(valCur[currency][:])

	valueCurrency, err := strconv.ParseFloat(srtingValueCur, 64)
	if err != nil {
		log.Printf("Error parcing Float")
	}
	return valueCurrency
}

func UsdAmount(usdToEur, rub, amount float64) float64 {
	usdToRub := rub / usdToEur
	usdAmount := amount / usdToRub
	return usdAmount
}
