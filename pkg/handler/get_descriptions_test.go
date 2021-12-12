package handler

// import (
// 	"bytes"
// 	"context"
// 	"net/http"
// 	"net/http/httptest"
// 	"tech_task/pkg/service"
// 	mock_service "tech_task/pkg/service/mocks"
// 	"testing"

// 	"github.com/go-chi/chi"
// 	"github.com/golang/mock/gomock"
// 	"gotest.tools/assert"
// )

// func TestHandler_GetAllUsersDescriptionsSort(t *testing.T) {
// 	type mockBehaviorGetDescriptions func(r *mock_service.MockGetDescriptions, sortBy, orderBy, sqlOrderBy string)
// 	ctx := context.Background()

// 	tests := []struct {
// 		name                        string
// 		inputBody                   string
// 		uid                         int64
// 		sortBy                      string
// 		orderBy                     string
// 		sqlOrderBy                  string
// 		mockBehaviorGetDescriptions mockBehaviorGetDescriptions
// 		expectedStatusCode          int
// 		expectedResponseBody        string
// 	}{
// 		{
// 			name:       "Ok",
// 			inputBody:  `{"id": "1","amount": "3680","description": "Покупка наушников","sender_receiver": "Avito","refill": "F"}`,
// 			uid:        20,
// 			sortBy:     "",
// 			orderBy:    "",
// 			sqlOrderBy: "ORDER BY",
// 			// mockBehaviorGetDescriptions: func(r *mock_service.MockGetDescriptions, id int64, sortBy, orderBy, sqlOrderBy string) {
// 			// 	r.EXPECT().GetDescriptionsUsers(ctx, id, sortBy, orderBy, sqlOrderBy).Return([]tech_task.Description{}, error)
// 			// },
// 			expectedStatusCode:   200,
// 			expectedResponseBody: "{\"user id\":1,\"balance at moment\":1839.55,\"amount\":3680,\"description of transaction\":\"Покупка наушников\",\"sender or receiver\":\"Avito\",\"refill the balance\":\"F\"}\n",
// },
// {
// 	name:           "Wrong input user",
// 	inputBody:      `{"id": "-1","amount": "3680","description": "Покупка наушников","sender_receiver": "Avito","refill": "F"}`,
// 	inputUser:      -1,
// 	balance:        10839.55,
// 	inputAmount:    3680,
// 	refill:         "F",
// 	description:    "Покупка наушников",
// 	senderReceiver: "Avito",
// 	mockBehaviorAddDescription: func(r *mock_service.MockAddDescription, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) {
// 	},
// 	mockBehaviorUpBalance:   func(r *mock_service.MockUpBalance, id int64, amount float64) {},
// 	mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {},
// 	mockBehaviorWritingOff:  func(r *mock_service.MockWritingOff, id int64, amount float64) {},
// 	expectedStatusCode:      400,
// 	expectedResponseBody:    "{\"error\":\"incorrect value id user\"}\n",
// },
// {
// 	name:           "Wrong input amount",
// 	inputBody:      `{"id": "1","amount": "-3680","description": "Покупка наушников","sender_receiver": "Avito","refill": "F"}`,
// 	inputUser:      1,
// 	balance:        10839.55,
// 	inputAmount:    -3680,
// 	refill:         "F",
// 	description:    "Покупка наушников",
// 	senderReceiver: "Avito",
// 	mockBehaviorAddDescription: func(r *mock_service.MockAddDescription, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) {
// 	},
// 	mockBehaviorUpBalance:   func(r *mock_service.MockUpBalance, id int64, amount float64) {},
// 	mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {},
// 	mockBehaviorWritingOff:  func(r *mock_service.MockWritingOff, id int64, amount float64) {},
// 	expectedStatusCode:      400,
// 	expectedResponseBody:    "{\"error\":\"the amount is negative\"}\n",
// },
// {
// 	name:           "Wrong input more 2 decimal places",
// 	inputBody:      `{"id": "1","amount": "3680.9876543","description": "Покупка наушников","sender_receiver": "Avito","refill": "F"}`,
// 	inputUser:      1,
// 	balance:        10839.55,
// 	inputAmount:    3680.9876543,
// 	refill:         "F",
// 	description:    "Покупка наушников",
// 	senderReceiver: "Avito",
// 	mockBehaviorAddDescription: func(r *mock_service.MockAddDescription, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) {
// 	},
// 	mockBehaviorUpBalance:   func(r *mock_service.MockUpBalance, id int64, amount float64) {},
// 	mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {},
// 	mockBehaviorWritingOff:  func(r *mock_service.MockWritingOff, id int64, amount float64) {},
// 	expectedStatusCode:      400,
// 	expectedResponseBody:    "{\"error\":\"the amount have more then 2 decimal places\"}\n",
// },
// {
// 	name:           "User not found",
// 	inputBody:      `{"id": "987654321","amount": "3680.98","description": "Покупка наушников","sender_receiver": "Avito","refill": "F"}`,
// 	inputUser:      987654321,
// 	balance:        10839.55,
// 	inputAmount:    3680.98,
// 	refill:         "F",
// 	description:    "Покупка наушников",
// 	senderReceiver: "Avito",
// 	mockBehaviorAddDescription: func(r *mock_service.MockAddDescription, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) {
// 	},
// 	mockBehaviorUpBalance: func(r *mock_service.MockUpBalance, id int64, amount float64) {},
// 	mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {
// 		var uid int64 = 0
// 		var balance float64 = 0
// 		var err error = errors.New("{\"error\":\"User not found\"}\n")
// 		r.EXPECT().BalanceInfoUser(ctx, id).Return(uid, balance, err)
// 	},
// 	mockBehaviorWritingOff: func(r *mock_service.MockWritingOff, id int64, amount float64) {},
// 	expectedStatusCode:     400,
// 	expectedResponseBody:   "{\"error\":\"User not found\"}\n",
// },
// {
// 	name:           "Wrong input refil",
// 	inputBody:      `{"id": "1","amount": "3680.98","description": "Покупка наушников","sender_receiver": "Avito","refill": ""}`,
// 	inputUser:      987654321,
// 	balance:        10839.55,
// 	inputAmount:    3680.98,
// 	refill:         "",
// 	description:    "Покупка наушников",
// 	senderReceiver: "Avito",
// 	mockBehaviorAddDescription: func(r *mock_service.MockAddDescription, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) {
// 	},
// 	mockBehaviorUpBalance:   func(r *mock_service.MockUpBalance, id int64, amount float64) {},
// 	mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {},
// 	mockBehaviorWritingOff:  func(r *mock_service.MockWritingOff, id int64, amount float64) {},
// 	expectedStatusCode:      400,
// 	expectedResponseBody:    "{\"error\":\"Refill is not null field\"}\n",
// },
//}
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			c := gomock.NewController(t)
// 			defer c.Finish()

// 			repo := mock_service.NewMocketGetDescriptions(c)
// 			test.mockBehaviorGetDescriptions(repo, test.uid, test.sortBy, test.orderBy, test.sqlOrderBy)

// 			services := &service.Service{
// 				GetDescriptions: repo,
// 			}
// 			handler := Handler{services}

// 			r := chi.NewRouter()
// 			r.Post("/description/add", handler.AddDescription)

// 			w := httptest.NewRecorder()
// 			req := httptest.NewRequest(http.MethodPost, "/description/add",
// 				bytes.NewBufferString(test.inputBody))

// 			r.ServeHTTP(w, req)

// 			assert.Equal(t, w.Code, test.expectedStatusCode)
// 			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
// 		})
// 	}
//
