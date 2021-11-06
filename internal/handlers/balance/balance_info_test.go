package balance

import (
	"bytes"
	"fmt"

	"github.com/stretchr/testify/mock"

	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockRepository struct {
	mock.Mock
}

type mockUser struct {
	Id      int64
	Balance float64
}


var mUser = mockUser{ 1,1059.55 }

func TestBalanceInfo(t *testing.T) {
	t.Run("good user_id", func(t *testing.T) {
		//JSONParams := bytes.NewBuffer([]byte(
		//	`{
		//	"id":"1"
		//	}`))

		//w := httptest.NewRecorder()
		//req, err := http.NewRequest("GET", "localhost:9000/balance-info", JSONParams)
		//if err != nil {
		//	t.Fatal(err)
		//}

		//req = req.WithContext(ctx)
		//repo := new(mockRepository)
		//repo.On("BalanceInfo", 1,1059.55).Return(mUser.Balance, nil)
		//res := w.Result()
		//defer res.Body.Close()
		//x, err := io.ReadAll(res.Body)
		//fmt.Println(repo)
		//fmt.Println(x)
		fmt.Println("rrr")


	})



	//
	//
	//r := httptest.NewRecorder()
	//handler := http.HandlerFunc(BalanceInfo)
	//handler.ServeHTTP(r, req)
	//
	//if status := r.Code; status != http.StatusOK {
	//	t.Errorf("handler returned wrong status code: got %v want %v",
	//		status, http.StatusOK)
	//}
	//
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//result := string(body)
	//expected := "{\"user id\":1,\"top up an amount\":1059.55}\n"
	//
	//repo.On("BalanceInfo", 1,1059.55).Return(mUser.mockUser, nil)
	//
	////service := HttpService{
	////	UserService: repo,
	////}
	////service.GetReceivers(w, r)
	//res := w.Result()
	//defer res.Body.Close()
	//str, err := io.ReadAll(res.Body)
	//require.Nil(t, err)
	//var resBody mockReceivers
	//err = json.Unmarshal(str, &resBody)
	//require.Nil(t, err)
	//require.Equal(t, http.StatusOK, w.Result().StatusCode)
	//require.Equal(t, mUser.Receivers, resBody.Receivers)
	//
	//if result != expected {
	//	t.Errorf("handler body: got %v want %s",
	//		result, expected)
	//}
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
