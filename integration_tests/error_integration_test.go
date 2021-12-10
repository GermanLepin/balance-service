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

const ServeAddress = "localhost:9000"

type ErrorSuite struct {
	suite.Suite
}

func TestErrorSuite(t *testing.T) {
	suite.Run(t, new(ErrorSuite))
}

func (b *ErrorSuite) TestIntegration_ErrorIntegration() {
	tests := []struct {
		name                 string
		inputBody            string
		url                  string
		http                 string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "ErrorUserId",
			inputBody:            `{"id":"-111111","amount":"1000.55"}`,
			http:                 http.MethodPost,
			url:                  fmt.Sprintf("http://%s/up-balance", ServeAddress),
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"incorrect value id user\"}\n",
		},
		{
			name:                 "ErrorAmount",
			inputBody:            `{"id":"111111","amount":"250.5556"}`,
			http:                 http.MethodPatch,
			url:                  fmt.Sprintf("http://%s/writing-off", ServeAddress),
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"the amount have more then 2 decimal places\"}\n",
		},
		{
			name:                 "ErrorNegativeAmount",
			inputBody:            `{"id":"111111","amount":"-250.55"}`,
			http:                 http.MethodPatch,
			url:                  fmt.Sprintf("http://%s/writing-off", ServeAddress),
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"the amount is negative\"}\n",
		},
		{
			name:                 "UserNotFound",
			inputBody:            `{"id":"99999999999"}`,
			http:                 http.MethodGet,
			url:                  fmt.Sprintf("http://%s/balance-info", ServeAddress),
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"User not found\"}\n",
		},
		{
			name:                 "BalanceInfoErrorConverUSDRequest",
			inputBody:            `{"id":"111111"}`,
			http:                 http.MethodGet,
			url:                  fmt.Sprintf("http://%s/balance-info?convert&currency=US", ServeAddress),
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"Invalid currency type, only USD\"}\n",
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
