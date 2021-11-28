package integration_tests

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

const ServeAddress = "localhost:9000"

type BalanceBillingSuite struct {
	suite.Suite
}

func TestBalanceBillingSuite(t *testing.T) {
	suite.Run(t, new(BalanceBillingSuite))
}

func (b *BalanceBillingSuite) TestUpBalanceRequest() {
	JSONParams := bytes.NewBuffer([]byte(
		`{
			"id":"111111",
			"amount":"1000.55"
		}`))

	balanceExpected := "{\"user id\":111111,\"top up an amount\":1000.55}\n"

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/up-balance", ServeAddress), JSONParams)
	b.NoError(err)

	client := http.Client{}
	result, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer result.Body.Close()

	upBalance, err := io.ReadAll(result.Body)
	b.NoError(err)

	b.Equal(http.StatusOK, result.StatusCode)
	b.Equal(balanceExpected, string(upBalance))
	b.NoError(err)
}

func (b *BalanceBillingSuite) TestWritingOffRequest() {
	JSONParams := bytes.NewBuffer([]byte(
		`{
			"id":"111111",
			"amount":"250.55"
		}`))

	balanceExpected := "{\"user id\":111111,\"writing off an amount\":250.55}\n"

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("http://%s/writing-off", ServeAddress), JSONParams)
	b.NoError(err)

	client := http.Client{}
	result, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer result.Body.Close()

	upBalance, err := io.ReadAll(result.Body)
	b.NoError(err)

	b.Equal(http.StatusOK, result.StatusCode)
	b.Equal(balanceExpected, string(upBalance))
	b.NoError(err)
}

func (b *BalanceBillingSuite) TestUserToUserRequest() {
	JSONParams := bytes.NewBuffer([]byte(`{
		"id1":"111111",
		"id2":"222222",
		"amount":"370"
	}`))
	balanceExpected := "{\"user id sender\":111111,\"writing off an amount\":370,\"user id recipient\":222222}\n"

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("http://%s/user-to-user", ServeAddress), JSONParams)
	b.NoError(err)

	client := http.Client{}
	result, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer result.Body.Close()

	balance, err := io.ReadAll(result.Body)
	b.NoError(err)

	b.Equal(http.StatusOK, result.StatusCode)
	b.Equal(balanceExpected, string(balance))
	b.NoError(err)
}

func (b *BalanceBillingSuite) TestBalanceInfoRequest() {
	JSONParams := bytes.NewBuffer([]byte(`{
		"id":"111111"
	}`))
	balanceExpected := "{\"user id\":111111,\"balance\":380}\n"

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/balance-info", ServeAddress), JSONParams)
	b.NoError(err)

	client := http.Client{}
	result, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer result.Body.Close()

	balance, err := io.ReadAll(result.Body)
	b.NoError(err)

	b.Equal(http.StatusOK, result.StatusCode)
	b.Equal(balanceExpected, string(balance))
	b.NoError(err)
}

func (b *BalanceBillingSuite) TestBalanceInfoConverUSDRequest() {
	JSONParams := bytes.NewBuffer([]byte(`{
		"id":"111111"
	}`))

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/balance-info?currency=USD", ServeAddress), JSONParams)
	b.NoError(err)

	client := http.Client{}
	result, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer result.Body.Close()

	balance, err := io.ReadAll(result.Body)
	b.NoError(err)

	b.Equal(http.StatusOK, result.StatusCode)
	b.NoError(err)
	fmt.Println(string(balance))
	fmt.Println("ok")
}
