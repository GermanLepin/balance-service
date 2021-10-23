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
	map_cur    map[string]json.RawMessage
	val_cur    map[string]json.RawMessage
	rates      = "rates"
	access_key = "27c4039d0e33e2f74fbdc7afa63c08a8"
)

func GetConvertValue(w http.ResponseWriter, currency string) float64 {
	client := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://api.exchangeratesapi.io/v1/latest?access_key=%s&symbols=%s", access_key, currency), nil)
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

	if err := json.Unmarshal(body, &map_cur); err != nil {
		log.Printf("Error parcing JSON")
	}

	map_currency := map_cur[rates]

	if err := json.Unmarshal(map_currency, &val_cur); err != nil {
		log.Printf("Error parcing JSON")
	}

	srting_value_cur := string(val_cur[currency][:])

	value_currency, err := strconv.ParseFloat(srting_value_cur, 64)
	if err != nil {
		log.Printf("Error parcing Float")
	}
	return value_currency
}

func UsdAmount(usd_to_eur, rub, amount float64) float64 {
	usd_to_rub := rub / usd_to_eur
	usd_amount := amount / usd_to_rub
	return usd_amount
}
