//go:build integration
// +build integration

package integration_test

// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"testing"

// 	"github.com/stretchr/testify/suite"
// )

// type ErrorSuite struct {
// 	suite.Suite
// }

// var ServeAddress = "localhost:9000"

// func TestErrorSuite(t *testing.T) {
// 	suite.Run(t, new(ErrorSuite))
// }

// func (e *ErrorSuite) TestIntegration_ErrorUserId() {
// 	JSONParams := bytes.NewBuffer([]byte(
// 		`{
// 			"id":"-111111",
// 			"amount":"1000.55"
// 		}`))

// 	balanceExpected := "{\"error\":\"Incorrect value id user\"}\n"

// 	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/up-balance", ServeAddress), JSONParams)
// 	e.NoError(err)

// 	client := http.Client{}
// 	result, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer result.Body.Close()

// 	balance, err := io.ReadAll(result.Body)
// 	e.NoError(err)

// 	e.Equal(http.StatusBadRequest, result.StatusCode)
// 	e.Equal(balanceExpected, string(balance))
// 	e.NoError(err)
// }

// func (e *ErrorSuite) TestIntegration_ErrorAmount() {
// 	JSONParams := bytes.NewBuffer([]byte(
// 		`{
// 			"id":"111111",
// 			"amount":"-1000.557865764758"
// 		}`))

// 	balanceExpected := "{\"error\":\"The amount have more then 2 decimal places\"}\n"

// 	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/up-balance", ServeAddress), JSONParams)
// 	e.NoError(err)

// 	client := http.Client{}
// 	result, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer result.Body.Close()

// 	balance, err := io.ReadAll(result.Body)
// 	e.NoError(err)

// 	e.Equal(http.StatusBadRequest, result.StatusCode)
// 	e.Equal(balanceExpected, string(balance))
// 	e.NoError(err)
// }

// func (e *ErrorSuite) TestIntegration_BalanceInfoErrorFindUserIdDB() {
// 	JSONParams := bytes.NewBuffer([]byte(
// 		`{
// 			"id":"111111111111",
// 			"amount":"1000.55"
// 		}`))

// 	balanceExpected := "{\"error\":\"User not found\"}\n"

// 	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/balance-info", ServeAddress), JSONParams)
// 	e.NoError(err)

// 	client := http.Client{}
// 	result, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer result.Body.Close()

// 	balance, err := io.ReadAll(result.Body)
// 	e.NoError(err)

// 	e.Equal(http.StatusBadRequest, result.StatusCode)
// 	e.Equal(balanceExpected, string(balance))
// 	e.NoError(err)
// }

// func (b *ErrorSuite) TestIntegration_BalanceInfoErrorConverUSDRequest() {
// 	JSONParams := bytes.NewBuffer([]byte(`{
// 		"id":"111111"
// 	}`))

// 	balanceExpected := "{\"error\":\"Invalid currency type, only USD\"}\n"

// 	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/balance-info?currency=U", ServeAddress), JSONParams)
// 	b.NoError(err)

// 	client := http.Client{}
// 	result, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer result.Body.Close()

// 	balance, err := io.ReadAll(result.Body)
// 	b.NoError(err)

// 	b.Equal(http.StatusBadRequest, result.StatusCode)
// 	b.NoError(err)
// 	b.Equal(balanceExpected, string(balance))
// }
