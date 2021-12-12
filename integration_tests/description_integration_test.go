//go:build integration
// +build integration

package integration_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DescriptionSuite struct {
	suite.Suite
}

func TestDescriptionSuite(t *testing.T) {
	suite.Run(t, new(DescriptionSuite))
}

func (b *DescriptionSuite) TestIntegration_AddDescription() {
	tests := []struct {
		name                 string
		inputBody            string
		url                  string
		http                 string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "AddDescriptionRefil",
			inputBody:            `{"id": "999999","amount": "6780","description": "Продажа наушников","sender_receiver": "Avito","refill": "T"}`,
			http:                 http.MethodPost,
			url:                  fmt.Sprintf("http://%s/description/add", ServeAddress),
			expectedStatusCode:   200,
			expectedResponseBody: "{\"user id\":999999,\"balance at moment\":6780,\"amount\":6780,\"description of transaction\":\"Продажа наушников\",\"sender or receiver\":\"Avito\",\"refill the balance\":\"T\"}\n",
		},
		{
			name:                 "AddDescriptionWritingOff",
			inputBody:            `{"id": "999999","amount": "5320","description": "Покупка куртки","sender_receiver": "Avito","refill": "F"}`,
			http:                 http.MethodPost,
			url:                  fmt.Sprintf("http://%s/description/add", ServeAddress),
			expectedStatusCode:   200,
			expectedResponseBody: "{\"user id\":999999,\"balance at moment\":1460,\"amount\":5320,\"description of transaction\":\"Покупка куртки\",\"sender or receiver\":\"Avito\",\"refill the balance\":\"F\"}\n",
		},
	}
	for _, test := range tests {
		b.Run(test.name, func() {
			req, err := http.NewRequest(test.http, test.url,
				bytes.NewBufferString(test.inputBody))
			b.NoError(err)

			client := http.Client{}
			result, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
			}
			defer result.Body.Close()

			testResponseBody, err := io.ReadAll(result.Body)
			b.NoError(err)

			b.Equal(result.StatusCode, test.expectedStatusCode)
			b.Equal(test.expectedResponseBody, string(testResponseBody))
			b.NoError(err)

		})
	}
}

func (b *DescriptionSuite) TestIntegration_GetAllDescriptions() {
	tests := []struct {
		name                 string
		inputBody            string
		url                  string
		http                 string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "GetAllDescription",
			inputBody:            "",
			http:                 http.MethodGet,
			url:                  fmt.Sprintf("http://%s/description/get-all", ServeAddress),
			expectedStatusCode:   200,
			expectedResponseBody: "{\"Id\":1,\"SenderReceiver\":\"Avito\",\"Amount\":6780,\"Description\":\"Продажа наушников\",\"BalanceAtMoment\":6780,\"UserId\":999999,\"CreatedAt\":\"2021-12-10T20:36:00.502123Z\",\"Refil\":\"T\"}\n{\"Id\":2,\"SenderReceiver\":\"Avito\",\"Amount\":5320,\"Description\":\"Покупка куртки\",\"BalanceAtMoment\":1460,\"UserId\":999999,\"CreatedAt\":\"2021-12-10T20:36:00.512675Z\",\"Refil\":\"F\"}\n",
		},
		{
			name:                 "GetAllDescriptionSortAmount",
			inputBody:            `{"sort_by": "amount"}`,
			http:                 http.MethodGet,
			url:                  fmt.Sprintf("http://%s/description/get-all", ServeAddress),
			expectedStatusCode:   200,
			expectedResponseBody: "{\"Id\":1,\"SenderReceiver\":\"Avito\",\"Amount\":6780,\"Description\":\"Продажа наушников\",\"BalanceAtMoment\":6780,\"UserId\":999999,\"CreatedAt\":\"2021-12-10T20:36:00.502123Z\",\"Refil\":\"T\"}\n{\"Id\":2,\"SenderReceiver\":\"Avito\",\"Amount\":5320,\"Description\":\"Покупка куртки\",\"BalanceAtMoment\":1460,\"UserId\":999999,\"CreatedAt\":\"2021-12-10T20:36:00.512675Z\",\"Refil\":\"F\"}\n",
		},
		{
			name:                 "GetAllDescriptionSortDescAmount",
			inputBody:            `{"sort_by": "amount", "order_by":"desc"}`,
			http:                 http.MethodGet,
			url:                  fmt.Sprintf("http://%s/description/get-all", ServeAddress),
			expectedStatusCode:   200,
			expectedResponseBody: "{\"Id\":2,\"SenderReceiver\":\"Avito\",\"Amount\":5320,\"Description\":\"Покупка куртки\",\"BalanceAtMoment\":1460,\"UserId\":999999,\"CreatedAt\":\"2021-12-10T20:36:00.512675Z\",\"Refil\":\"F\"}\n{\"Id\":1,\"SenderReceiver\":\"Avito\",\"Amount\":6780,\"Description\":\"Продажа наушников\",\"BalanceAtMoment\":6780,\"UserId\":999999,\"CreatedAt\":\"2021-12-10T20:36:00.502123Z\",\"Refil\":\"T\"}\n",
		},
	}
	for _, test := range tests {
		b.Run(test.name, func() {
			req, err := http.NewRequest(test.http, test.url,
				bytes.NewBufferString(test.inputBody))
			b.NoError(err)

			client := http.Client{}
			result, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
			}
			defer result.Body.Close()

			b.Equal(result.StatusCode, test.expectedStatusCode)
			b.NoError(err)
		})
	}
}

func (b *DescriptionSuite) TestIntegration_GetUserIdDescriptions() {
	tests := []struct {
		name                 string
		inputBody            string
		url                  string
		http                 string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "GetUserIdDescription",
			inputBody:            `{"id": "999999"}`,
			http:                 http.MethodGet,
			url:                  fmt.Sprintf("http://%s/description/get-user", ServeAddress),
			expectedStatusCode:   200,
			expectedResponseBody: "{\"Id\":1,\"SenderReceiver\":\"Avito\",\"Amount\":6780,\"Description\":\"Продажа наушников\",\"BalanceAtMoment\":6780,\"UserId\":999999,\"CreatedAt\":\"2021-12-10T20:36:00.502123Z\",\"Refil\":\"T\"}\n",
		},
		{
			name:                 "GetUserIdDescriptionSortAmount",
			inputBody:            `{"id": "999999","sort_by": "amount"}`,
			http:                 http.MethodGet,
			url:                  fmt.Sprintf("http://%s/description/get-user", ServeAddress),
			expectedStatusCode:   200,
			expectedResponseBody: "{\"Id\":1,\"SenderReceiver\":\"Avito\",\"Amount\":6780,\"Description\":\"Продажа наушников\",\"BalanceAtMoment\":6780,\"UserId\":999999,\"CreatedAt\":\"2021-12-10T20:36:00.502123Z\",\"Refil\":\"T\"}\n",
		},
		{
			name:                 "GetUserIdDescriptionSortDescAmount",
			inputBody:            `{"id": "999999","sort_by": "amount", "order_by":"desc"}`,
			http:                 http.MethodGet,
			url:                  fmt.Sprintf("http://%s/description/get-user", ServeAddress),
			expectedStatusCode:   200,
			expectedResponseBody: "{\"Id\":1,\"SenderReceiver\":\"Avito\",\"Amount\":6780,\"Description\":\"Продажа наушников\",\"BalanceAtMoment\":6780,\"UserId\":999999,\"CreatedAt\":\"2021-12-10T20:36:00.502123Z\",\"Refil\":\"T\"}\n",
		},
	}
	for _, test := range tests {
		b.Run(test.name, func() {
			req, err := http.NewRequest(test.http, test.url,
				bytes.NewBufferString(test.inputBody))
			b.NoError(err)

			client := http.Client{}
			result, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
			}
			defer result.Body.Close()

			b.Equal(result.StatusCode, test.expectedStatusCode)
			b.NoError(err)
		})
	}
}
