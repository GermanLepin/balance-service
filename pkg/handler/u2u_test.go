package handler

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"tech_task/pkg/service"
	mock_service "tech_task/pkg/service/mocks"
	"testing"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestHandler_U2U(t *testing.T) {
	type mockBehaviorUpBalance func(r *mock_service.MockUpBalance, id int64, amount float64)
	type mockBehaviorBalanceInfo func(r *mock_service.MockBalanceInfo, id int64)
	type mockBehaviorWritingOff func(r *mock_service.MockWritingOff, id int64, amount float64)
	ctx := context.Background()

	tests := []struct {
		name                    string
		inputBody               string
		inputFirstUser          int64
		inputSecondUser         int64
		inputAmount             float64
		mockBehaviorUpBalance   mockBehaviorUpBalance
		mockBehaviorBalanceInfo mockBehaviorBalanceInfo
		mockBehaviorWritingOff  mockBehaviorWritingOff
		expectedStatusCode      int
		expectedResponseBody    string
	}{
		{
			name:            "Ok",
			inputBody:       `{"user id1":"1","user id2":"2","amount":"743.63"}`,
			inputFirstUser:  1,
			inputSecondUser: 2,
			inputAmount:     743.63,
			mockBehaviorUpBalance: func(r *mock_service.MockUpBalance, id int64, amount float64) {
				var err error = nil
				r.EXPECT().UpBalanceUser(ctx, id, amount).Return(err)
			},
			mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {
				var uid int64 = 1
				var balance = 1839.55
				var err error = nil
				r.EXPECT().BalanceInfoUser(ctx, id).Return(uid, balance, err)
			},
			mockBehaviorWritingOff: func(r *mock_service.MockWritingOff, id int64, amount float64) {
				var uid int64 = 1
				var respAmount = 743.63
				var err error = nil
				r.EXPECT().WritingOffUser(ctx, id, amount).Return(uid, respAmount, err)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"user id sender\":1,\"writing off an amount\":743.63,\"user id recipient\":2}\n",
		},
		{
			name:                    "Wrong input user",
			inputBody:               `{"user id1":"-1","user id2":"2","amount":"743.63"}`,
			inputFirstUser:          -1,
			inputSecondUser:         2,
			inputAmount:             1569.77,
			mockBehaviorUpBalance:   func(r *mock_service.MockUpBalance, id int64, amount float64) {},
			mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {},
			mockBehaviorWritingOff:  func(r *mock_service.MockWritingOff, id int64, amount float64) {},
			expectedStatusCode:      400,
			expectedResponseBody:    "{\"error\":\"incorrect value user id\"}\n",
		},
		{
			name:                    "Wrong input amount",
			inputBody:               `{"user id1":"1","user id2":"2","amount":"-743.63"}`,
			inputFirstUser:          1,
			inputSecondUser:         2,
			inputAmount:             -743.63,
			mockBehaviorUpBalance:   func(r *mock_service.MockUpBalance, id int64, amount float64) {},
			mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {},
			mockBehaviorWritingOff:  func(r *mock_service.MockWritingOff, id int64, amount float64) {},
			expectedStatusCode:      400,
			expectedResponseBody:    "{\"error\":\"the amount is negative\"}\n",
		},
		{
			name:                    "Wrong input more 2 decimal places",
			inputBody:               `{"user id1":"1","user id2":"2","amount":"743.63453"}`,
			inputFirstUser:          1,
			inputSecondUser:         2,
			inputAmount:             743.63453,
			mockBehaviorUpBalance:   func(r *mock_service.MockUpBalance, id int64, amount float64) {},
			mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {},
			mockBehaviorWritingOff:  func(r *mock_service.MockWritingOff, id int64, amount float64) {},
			expectedStatusCode:      400,
			expectedResponseBody:    "{\"error\":\"the amount have more then 2 decimal places\"}\n",
		},
		{
			name:                  "User not found",
			inputBody:             `{"user id1":"999888","user id2":"2","amount":"743.63"}`,
			inputFirstUser:        999888,
			inputSecondUser:       2,
			inputAmount:           743.63,
			mockBehaviorUpBalance: func(r *mock_service.MockUpBalance, id int64, amount float64) {},
			mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {
				var uid int64 = 0
				var balance float64 = 0
				var err = errors.New("{\"error\":\"User not found\"}\n")
				r.EXPECT().BalanceInfoUser(ctx, id).Return(uid, balance, err)
			},
			mockBehaviorWritingOff: func(r *mock_service.MockWritingOff, id int64, amount float64) {},
			expectedStatusCode:     400,
			expectedResponseBody:   "{\"error\":\"User not found\"}\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repoInfo := mock_service.NewMockBalanceInfoUser(c)
			test.mockBehaviorBalanceInfo(repoInfo, test.inputFirstUser)

			repoWritingOff := mock_service.NewMockWritingOffUser(c)
			test.mockBehaviorWritingOff(repoWritingOff, test.inputFirstUser, test.inputAmount)

			repoUpBalance := mock_service.NewMockUpBalanceUser(c)
			test.mockBehaviorUpBalance(repoUpBalance, test.inputSecondUser, test.inputAmount)

			services := &service.Service{
				BalanceInfo: repoInfo,
				WritingOff:  repoWritingOff,
				UpBalance:   repoUpBalance,
			}
			handler := Handler{services}

			r := chi.NewRouter()
			r.Patch("/user-to-user", handler.U2U)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPatch, "/user-to-user",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
