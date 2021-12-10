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

func TestHandler_WritingOff(t *testing.T) {
	type mockBehavior func(r *mock_service.MockWritingOff, id int64, amount float64)
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
			inputBody:   `{"id":"1","amount":"969.63"}`,
			inputUser:   1,
			inputAmount: 969.63,
			mockBehavior: func(r *mock_service.MockWritingOff, id int64, amount float64) {
				var uid int64 = 1
				var respAmount float64 = 969.63
				var err error = nil
				r.EXPECT().WritingOffUser(ctx, id, amount).Return(uid, respAmount, err)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"user id\":1,\"writing off an amount\":969.63}\n",
		},
		{
			name:                 "Wrong input user",
			inputBody:            `{"id":"-1","amount":"1569.77"}`,
			inputUser:            -1,
			inputAmount:          1569.77,
			mockBehavior:         func(r *mock_service.MockWritingOff, id int64, amount float64) {},
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"incorrect value id user\"}\n",
		},
		{
			name:                 "Wrong input amount",
			inputBody:            `{"id":"1","amount":"-1569.77"}`,
			inputUser:            1,
			inputAmount:          -1569.77,
			mockBehavior:         func(r *mock_service.MockWritingOff, id int64, amount float64) {},
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"the amount is negative\"}\n",
		},
		{
			name:                 "Wrong input more 2 decimal places",
			inputBody:            `{"id":"1","amount":"1569.77345"}`,
			inputUser:            1,
			inputAmount:          1569.77345,
			mockBehavior:         func(r *mock_service.MockWritingOff, id int64, amount float64) {},
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"the amount have more then 2 decimal places\"}\n",
		},
		{
			name:                 "User not found",
			inputBody:            `{"id":"1","amount":"1569.77"}`,
			inputUser:            99999999,
			inputAmount:          1569.77,
			mockBehavior:         func(r *mock_service.MockWritingOff, id int64, amount float64) {},
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"user not found\"}\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockWritingOffUser(c)
			test.mockBehavior(repo, test.inputUser, test.inputAmount)

			servicesWritingOff := &service.Service{WritingOff: repo}
			handler := Handler{servicesWritingOff}

			r := chi.NewRouter()
			r.Patch("/writing-off", handler.WritingOff)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPatch, "/writing-off",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
