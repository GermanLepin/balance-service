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

func TestHandler_AddDescription(t *testing.T) {
	type mockBehaviorBalanceInfo func(r *mock_service.MockBalanceInfo, id int64)
	type mockBehaviorWritingOff func(r *mock_service.MockWritingOff, id int64, amount float64)
	type mockBehaviorUpBalance func(r *mock_service.MockUpBalance, id int64, amount float64)
	type mockBehaviorAddDescription func(r *mock_service.MockAddDescription, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string)
	ctx := context.Background()

	tests := []struct {
		name                       string
		inputBody                  string
		inputUser                  int64
		balance                    float64
		inputAmount                float64
		refill                     string
		description                string
		senderReceiver             string
		mockBehaviorUpBalance      mockBehaviorUpBalance
		mockBehaviorBalanceInfo    mockBehaviorBalanceInfo
		mockBehaviorWritingOff     mockBehaviorWritingOff
		mockBehaviorAddDescription mockBehaviorAddDescription
		expectedStatusCode         int
		expectedResponseBody       string
	}{
		{
			name:           "Add desciption refill true",
			inputBody:      `{"user id":"1","amount": "3680","description": "Покупка наушников","sender receiver": "Avito","refill": "T"}`,
			inputUser:      1,
			balance:        1085.55,
			inputAmount:    3680,
			refill:         "T",
			description:    "Покупка наушников",
			senderReceiver: "Avito",
			mockBehaviorUpBalance: func(r *mock_service.MockUpBalance, id int64, amount float64) {
				var err error = nil
				r.EXPECT().UpBalanceUser(ctx, id, amount).Return(err)
			},
			mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {
				var uid int64 = 1
				var balance = 1085.55
				var err error = nil
				r.EXPECT().BalanceInfoUser(ctx, id).Return(uid, balance, err)
			},
			mockBehaviorWritingOff: func(r *mock_service.MockWritingOff, id int64, amount float64) {},
			mockBehaviorAddDescription: func(r *mock_service.MockAddDescription, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) {
				var err error = nil
				r.EXPECT().AddDescriptionUser(ctx, id, balanceAtMoment, corectAmount, refill, description, senderReceiver).Return(err)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"user id\":1,\"balance at moment\":1085.55,\"amount\":3680,\"description of transaction\":\"Покупка наушников\",\"sender or receiver\":\"Avito\",\"refill the balance\":\"T\"}\n",
		},
		{
			name:           "Wrong input user",
			inputBody:      `{"user id": "-1","amount": "3680","description": "Покупка наушников","sender receiver": "Avito","refill": "F"}`,
			inputUser:      -1,
			balance:        10839.55,
			inputAmount:    3680,
			refill:         "F",
			description:    "Покупка наушников",
			senderReceiver: "Avito",
			mockBehaviorAddDescription: func(r *mock_service.MockAddDescription, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) {
			},
			mockBehaviorUpBalance:   func(r *mock_service.MockUpBalance, id int64, amount float64) {},
			mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {},
			mockBehaviorWritingOff:  func(r *mock_service.MockWritingOff, id int64, amount float64) {},
			expectedStatusCode:      400,
			expectedResponseBody:    "{\"error\":\"incorrect value user id\"}\n",
		},
		{
			name:           "Wrong input amount",
			inputBody:      `{"user id": "1","amount": "-3680","description": "Покупка наушников","sender receiver": "Avito","refill": "F"}`,
			inputUser:      1,
			balance:        10839.55,
			inputAmount:    -3680,
			refill:         "F",
			description:    "Покупка наушников",
			senderReceiver: "Avito",
			mockBehaviorAddDescription: func(r *mock_service.MockAddDescription, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) {
			},
			mockBehaviorUpBalance:   func(r *mock_service.MockUpBalance, id int64, amount float64) {},
			mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {},
			mockBehaviorWritingOff:  func(r *mock_service.MockWritingOff, id int64, amount float64) {},
			expectedStatusCode:      400,
			expectedResponseBody:    "{\"error\":\"the amount is negative\"}\n",
		},
		{
			name:           "Wrong input more 2 decimal places",
			inputBody:      `{"user id": "1","amount": "3680.9876543","description": "Покупка наушников","sender receiver": "Avito","refill": "F"}`,
			inputUser:      1,
			balance:        10839.55,
			inputAmount:    3680.9876543,
			refill:         "F",
			description:    "Покупка наушников",
			senderReceiver: "Avito",
			mockBehaviorAddDescription: func(r *mock_service.MockAddDescription, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) {
			},
			mockBehaviorUpBalance:   func(r *mock_service.MockUpBalance, id int64, amount float64) {},
			mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {},
			mockBehaviorWritingOff:  func(r *mock_service.MockWritingOff, id int64, amount float64) {},
			expectedStatusCode:      400,
			expectedResponseBody:    "{\"error\":\"the amount have more then 2 decimal places\"}\n",
		},
		{
			name:           "User not found",
			inputBody:      `{"user id": "987654321","amount": "3680.98","description": "Покупка наушников","sender receiver": "Avito","refill": "F"}`,
			inputUser:      987654321,
			balance:        10839.55,
			inputAmount:    3680.98,
			refill:         "F",
			description:    "Покупка наушников",
			senderReceiver: "Avito",
			mockBehaviorAddDescription: func(r *mock_service.MockAddDescription, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) {
			},
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
		{
			name:           "Wrong input refill",
			inputBody:      `{"user id": "1","amount": "3680.98","description": "Покупка наушников","sender receiver": "Avito","refill": ""}`,
			inputUser:      987654321,
			balance:        10839.55,
			inputAmount:    3680.98,
			refill:         "",
			description:    "Покупка наушников",
			senderReceiver: "Avito",
			mockBehaviorAddDescription: func(r *mock_service.MockAddDescription, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) {
			},
			mockBehaviorUpBalance:   func(r *mock_service.MockUpBalance, id int64, amount float64) {},
			mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {},
			mockBehaviorWritingOff:  func(r *mock_service.MockWritingOff, id int64, amount float64) {},
			expectedStatusCode:      400,
			expectedResponseBody:    "{\"error\":\"Refill is not null field\"}\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repoUpBalance := mock_service.NewMockUpBalanceUser(c)
			test.mockBehaviorUpBalance(repoUpBalance, test.inputUser, test.inputAmount)

			repoInfo := mock_service.NewMockBalanceInfoUser(c)
			test.mockBehaviorBalanceInfo(repoInfo, test.inputUser)

			repoWritingOff := mock_service.NewMockWritingOffUser(c)
			test.mockBehaviorWritingOff(repoWritingOff, test.inputUser, test.inputAmount)

			repoAddDescription := mock_service.NewMockAddDescriptionUser(c)
			test.mockBehaviorAddDescription(repoAddDescription, test.inputUser, test.balance, test.inputAmount, test.refill, test.description, test.senderReceiver)

			services := &service.Service{
				UpBalance:      repoUpBalance,
				BalanceInfo:    repoInfo,
				WritingOff:     repoWritingOff,
				AddDescription: repoAddDescription,
			}
			handler := Handler{services}

			r := chi.NewRouter()
			r.Post("/description/add", handler.AddDescription)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/description/add",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
