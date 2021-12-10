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

type BalanceBillingSuite struct {
	suite.Suite
}

func TestBalanceBillingSuite(t *testing.T) {
	suite.Run(t, new(BalanceBillingSuite))
}

func (b *BalanceBillingSuite) TestIntegration_BalanceBiling() {
	tests := []struct {
		name                 string
		inputBody            string
		url                  string
		http                 string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "UpBalanceRequest",
			inputBody:            `{"id":"111111","amount":"1000.55"}`,
			http:                 http.MethodPost,
			url:                  fmt.Sprintf("http://%s/up-balance", ServeAddress),
			expectedStatusCode:   200,
			expectedResponseBody: "{\"user id\":111111,\"top up an amount\":1000.55}\n",
		},
		{
			name:                 "WritingOffRequest",
			inputBody:            `{"id":"111111","amount":"250.55"}`,
			http:                 http.MethodPatch,
			url:                  fmt.Sprintf("http://%s/writing-off", ServeAddress),
			expectedStatusCode:   200,
			expectedResponseBody: "{\"user id\":111111,\"writing off an amount\":250.55}\n",
		},
		{
			name:                 "UserToUserRequest",
			inputBody:            `{"id1":"111111","id2":"222222","amount":"349.99"}`,
			http:                 http.MethodPatch,
			url:                  fmt.Sprintf("http://%s/user-to-user", ServeAddress),
			expectedStatusCode:   200,
			expectedResponseBody: "{\"user id sender\":111111,\"writing off an amount\":349.99,\"user id recipient\":222222}\n",
		},
		{
			name:                 "BalanceInfoRequest",
			inputBody:            `{"id":"111111"}`,
			http:                 http.MethodGet,
			url:                  fmt.Sprintf("http://%s/balance-info", ServeAddress),
			expectedStatusCode:   200,
			expectedResponseBody: "{\"user id\":111111,\"balance\":400.01}\n",
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
