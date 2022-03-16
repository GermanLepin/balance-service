package handler

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"tech_task/pkg/service"
	mock_service "tech_task/pkg/service/mocks"
	"testing"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestHandler_UpBalance(t *testing.T) {
	type mockBehavior func(r *mock_service.MockUpBalance, id int64, amount float64)
	ctx := context.Background()

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            int64
		inputAmount          float64
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Ok",
			inputBody:   `{"user id":"1","amount":"1569.77"}`,
			inputUser:   1,
			inputAmount: 1569.77,
			mockBehavior: func(r *mock_service.MockUpBalance, id int64, amount float64) {
				var err error = nil
				r.EXPECT().UpBalanceUser(ctx, id, amount).Return(err)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"user id\":1,\"top up an amount\":1569.77}\n",
		},
		{
			name:                 "Wrong input user",
			inputBody:            `{"user id":"-1","amount":"1569.77"}`,
			inputUser:            -1,
			inputAmount:          1569.77,
			mockBehavior:         func(r *mock_service.MockUpBalance, id int64, amount float64) {},
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"incorrect value user id\"}\n",
		},
		{
			name:                 "Wrong input amount",
			inputBody:            `{"user id":"1","amount":"-1569.77"}`,
			inputUser:            1,
			inputAmount:          -1569.77,
			mockBehavior:         func(r *mock_service.MockUpBalance, id int64, amount float64) {},
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"the amount is negative\"}\n",
		},
		{
			name:                 "Wrong input more 2 decimal places",
			inputBody:            `{"user id":"1","amount":"1569.77345"}`,
			inputUser:            1,
			inputAmount:          1569.77345,
			mockBehavior:         func(r *mock_service.MockUpBalance, id int64, amount float64) {},
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"the amount have more then 2 decimal places\"}\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockUpBalanceUser(c)
			test.mockBehavior(repo, test.inputUser, test.inputAmount)

			services := &service.Service{UpBalance: repo}
			handler := Handler{services}

			r := chi.NewRouter()
			r.Post("/up-balance", handler.UpBalance)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/up-balance",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
