package balance

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBalanceInfo(t *testing.T) {
	JSONparams := bytes.NewBuffer([]byte(
		`{
			"id":"1"
			}`))

	req, err := http.NewRequest("GET", "localhost:9000/balance-info", JSONparams)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(BalanceInfo)
	handler.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	result := string(body)
	expected := "{\"user id\":1,\"top up an amount\":1059.55}\n"

	if result != expected {
		t.Errorf("handler body: got %v want %s",
			result, expected)
	}
}

func TestUpBalanceErrorUserId(t *testing.T) {
	JSONparams := bytes.NewBuffer([]byte(
		`{
			"id":"-1",
			"amount":"10590.55"
		}`))

	req, err := http.NewRequest("POST", "localhost:9000/balance-info", JSONparams)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(BalanceInfo)
	handler.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	result := string(body)
	expected := "{\"error\":\"Incorrect value id user\"}\n"

	if result != expected {
		t.Errorf("handler body: got %v want %v",
			result, expected)
	}
}

func TestUpBalanceErrorAmount(t *testing.T) {
	JSONparams := bytes.NewBuffer([]byte(
		`{
			"id":"1",
			"amount":"-10590.55"
		}`))

	req, err := http.NewRequest("POST", "localhost:9000/balance-info", JSONparams)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(BalanceInfo)
	handler.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	result := string(body)
	expected := "{\"error\":\"The amount is negative\"}\n"

	if result != expected {
		t.Errorf("handler body: got %v want %v",
			result, expected)
	}
}

func TestBalanceInfoErrorFindUserIdDB(t *testing.T) {
	JSONparams := bytes.NewBuffer([]byte(
		`{
			"id":"111111111111111",
			"amount":"10590.55"
		}`))

	req, err := http.NewRequest("POST", "localhost:9000/balance-info", JSONparams)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(BalanceInfo)
	handler.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	result := string(body)
	expected := "{\"error\":\"User not found in database\"}\n"

	if result != expected {
		t.Errorf("handler body: got %v want %v",
			result, expected)
	}
}
