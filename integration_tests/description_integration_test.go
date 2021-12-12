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
			url:                  fmt.Sprintf("http://%s/descriptions/add", ServeAddress),
			expectedStatusCode:   200,
			expectedResponseBody: "{\"user id\":999999,\"balance at moment\":6780,\"amount\":6780,\"description of transaction\":\"Продажа наушников\",\"sender or receiver\":\"Avito\",\"refill the balance\":\"T\"}\n",
		},
		{
			name:                 "AddDescriptionWritingOff",
			inputBody:            `{"id": "999999","amount": "5320","description": "Покупка куртки","sender_receiver": "Avito","refill": "F"}`,
			http:                 http.MethodPost,
			url:                  fmt.Sprintf("http://%s/descriptions/add", ServeAddress),
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
		name               string
		inputBody          string
		url                string
		http               string
		expectedStatusCode int
	}{
		{
			name:               "GetAllDescription",
			inputBody:          "",
			http:               http.MethodGet,
			url:                fmt.Sprintf("http://%s/descriptions/get", ServeAddress),
			expectedStatusCode: 200,
		},
		{
			name:               "GetAllDescriptionSortAmount",
			inputBody:          `{"sort_by": "amount"}`,
			http:               http.MethodGet,
			url:                fmt.Sprintf("http://%s/descriptions/get", ServeAddress),
			expectedStatusCode: 200,
		},
		{
			name:               "GetAllDescriptionSortDescAmount",
			inputBody:          `{"sort_by": "amount", "order_by":"desc"}`,
			http:               http.MethodGet,
			url:                fmt.Sprintf("http://%s/descriptions/get", ServeAddress),
			expectedStatusCode: 200,
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
		name               string
		inputBody          string
		url                string
		http               string
		expectedStatusCode int
	}{
		{
			name:               "GetUserIdDescription",
			inputBody:          `{"id": "999999"}`,
			http:               http.MethodGet,
			url:                fmt.Sprintf("http://%s/descriptions/get", ServeAddress),
			expectedStatusCode: 200,
		},
		{
			name:               "GetUserIdDescriptionSortAmount",
			inputBody:          `{"id": "999999","sort_by": "amount"}`,
			http:               http.MethodGet,
			url:                fmt.Sprintf("http://%s/descriptions/get", ServeAddress),
			expectedStatusCode: 200,
		},
		{
			name:               "GetUserIdDescriptionSortDescAmount",
			inputBody:          `{"id": "999999","sort_by": "amount", "order_by":"desc"}`,
			http:               http.MethodGet,
			url:                fmt.Sprintf("http://%s/descriptions/get", ServeAddress),
			expectedStatusCode: 200,
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
