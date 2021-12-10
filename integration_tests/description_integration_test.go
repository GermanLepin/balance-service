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

// type DescriptionSuite struct {
// 	suite.Suite
// }

// func TestDescriptionSuite(t *testing.T) {
// 	suite.Run(t, new(DescriptionSuite))
// }

// func (d *DescriptionSuite) TestIntegration_AddDescriptionRefillTrue() {
// 	JSONParams := bytes.NewBuffer([]byte(
// 		`{
// 			"id": "999999",
// 			"amount": "6780",
// 			"description": "Продажа наушников",
// 			"sender_receiver": "Avito",
// 			"refill": "T"
// 		}`))

// 	balanceExpected := "{\"user id\":999999,\"balance at moment\":6780,\"amount\":6780,\"description of transaction\":\"Продажа наушников\",\"sender or receiver\":\"Avito\",\"refill the balance\":\"T\"}\n"

// 	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/description/add", ServeAddress), JSONParams)
// 	d.NoError(err)

// 	client := http.Client{}
// 	result, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer result.Body.Close()

// 	upBalance, err := io.ReadAll(result.Body)
// 	d.NoError(err)

// 	d.Equal(http.StatusOK, result.StatusCode)
// 	d.Equal(balanceExpected, string(upBalance))
// 	d.NoError(err)
// }

// func (d *DescriptionSuite) TestAddDescriptionRefillFalse() {
// 	JSONParams := bytes.NewBuffer([]byte(
// 		`{
// 			"id": "999999",
// 			"amount": "5320",
// 			"description": "Покупка куртки",
// 			"sender_receiver": "Avito",
// 			"refill": "F"
// 		}`))

// 	balanceExpected := "{\"user id\":999999,\"balance at moment\":1460,\"amount\":5320,\"description of transaction\":\"Покупка куртки\",\"sender or receiver\":\"Avito\",\"refill the balance\":\"F\"}\n"

// 	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/description/add", ServeAddress), JSONParams)
// 	d.NoError(err)

// 	client := http.Client{}
// 	result, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer result.Body.Close()

// 	upBalance, err := io.ReadAll(result.Body)
// 	d.NoError(err)

// 	d.Equal(http.StatusOK, result.StatusCode)
// 	d.Equal(balanceExpected, string(upBalance))
// 	d.NoError(err)
// }
