package handlers_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"tech_task/internal/handlers"
	"testing"
)

func TestU2U(t *testing.T) {
	JSONparams := bytes.NewBuffer([]byte(
		`{
			"id1":"1",
			"id2":"2",
			"amount":"60"
			}`))

	req := httptest.NewRequest("POST", "localhost:9000/user-to-user", JSONparams)
	r := httptest.NewRecorder()
	handlers.U2U(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	result := string(body)
	expected := "{\"user id sender\":1,\"writing off an amount\":60,\"user id recipient\":2}\n"

	if result != expected {
		t.Errorf("handler body: got %v want %s",
			result, expected)
	}
}

func TestU2UErrorUserId(t *testing.T) {
	JSONparams := bytes.NewBuffer([]byte(
		`{
			"id1":"1",
			"id2":"-2",
			"amount":"10590.55"
		}`))

	req := httptest.NewRequest("POST", "localhost:9000/user-to-user", JSONparams)
	r := httptest.NewRecorder()
	handlers.U2U(r, req)

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

func TestU2UErrorAmount(t *testing.T) {
	JSONparams := bytes.NewBuffer([]byte(
		`{
			"id1":"1",
			"id2":"2",
			"amount":"-10590.55"
		}`))

	req := httptest.NewRequest("POST", "localhost:9000/user-to-user", JSONparams)
	r := httptest.NewRecorder()
	handlers.U2U(r, req)

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

func TestU2UErrorFindUserIdDB(t *testing.T) {
	JSONparams := bytes.NewBuffer([]byte(
		`{
			"id1":"1111111111",
			"id2":"1",
			"amount":"10590.55"
		}`))

	req := httptest.NewRequest("POST", "localhost:9000/user-to-user", JSONparams)
	r := httptest.NewRecorder()
	handlers.U2U(r, req)

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
