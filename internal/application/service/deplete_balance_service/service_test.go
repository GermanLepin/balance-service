package deplete_balance_service

// import (
// 	"bytes"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"tech_task/pkg/service"
// 	mock_service "tech_task/pkg/service/mocks"
// 	"testing"

// 	"github.com/go-chi/chi"
// 	"github.com/golang/mock/gomock"
// 	"gotest.tools/assert"
// )

// func TestHandler_WritingOff(t *testing.T) {
// 	type mockBehaviorBalanceInfo func(r *mock_service.MockBalanceInfo, id int64)
// 	type mockBehaviorWritingOff func(r *mock_service.MockWritingOff, id int64, amount float32)

// 	tests := []struct {
// 		name                    string
// 		inputBody               string
// 		inputUser               int64
// 		inputAmount             float32
// 		mockBehaviorBalanceInfo mockBehaviorBalanceInfo
// 		mockBehaviorWritingOff  mockBehaviorWritingOff
// 		expectedStatusCode      int
// 		expectedResponseBody    string
// 	}{
// 		{
// 			name:        "Ok",
// 			inputBody:   `{"user id":"1","amount":"969.63"}`,
// 			inputUser:   1,
// 			inputAmount: 969.63,
// 			mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {
// 				var uid int64 = 1
// 				var balance = 1830.55
// 				var err error = nil
// 				r.EXPECT().BalanceInfoUser(ctx, id).Return(uid, balance, err)
// 			},
// 			mockBehaviorWritingOff: func(r *mock_service.MockWritingOff, id int64, amount float32) {
// 				var uid int64 = 1
// 				var respAmount = 969.63
// 				var err error = nil
// 				r.EXPECT().WritingOffUser(ctx, id, amount).Return(uid, respAmount, err)
// 			},
// 			expectedStatusCode:   200,
// 			expectedResponseBody: "{\"user id\":1,\"writing off an amount\":969.63}\n",
// 		},
// 		{
// 			name:                    "Wrong input user",
// 			inputBody:               `{"user id":"-1","amount":"1569.77"}`,
// 			inputUser:               -1,
// 			inputAmount:             1569.77,
// 			mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {},
// 			mockBehaviorWritingOff:  func(r *mock_service.MockWritingOff, id int64, amount float32) {},
// 			expectedStatusCode:      400,
// 			expectedResponseBody:    "{\"error\":\"incorrect value user id\"}\n",
// 		},
// 		{
// 			name:                    "Wrong input amount",
// 			inputBody:               `{"user id":"1","amount":"-1569.77"}`,
// 			inputUser:               1,
// 			inputAmount:             -1569.77,
// 			mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {},
// 			mockBehaviorWritingOff:  func(r *mock_service.MockWritingOff, id int64, amount float32) {},
// 			expectedStatusCode:      400,
// 			expectedResponseBody:    "{\"error\":\"the amount is negative\"}\n",
// 		},
// 		{
// 			name:                    "Wrong input more 2 decimal places",
// 			inputBody:               `{"user id":"1","amount":"1569.77345"}`,
// 			inputUser:               1,
// 			inputAmount:             1569.77345,
// 			mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {},
// 			mockBehaviorWritingOff:  func(r *mock_service.MockWritingOff, id int64, amount float32) {},
// 			expectedStatusCode:      400,
// 			expectedResponseBody:    "{\"error\":\"the amount have more then 2 decimal places\"}\n",
// 		},
// 		{
// 			name:        "User not found",
// 			inputBody:   `{"user id":"99999999","amount":"1569.77"}`,
// 			inputUser:   99999999,
// 			inputAmount: 1569.77,
// 			mockBehaviorBalanceInfo: func(r *mock_service.MockBalanceInfo, id int64) {
// 				var uid int64 = 0
// 				var balance float32 = 0
// 				var err = errors.New("{\"error\":\"User not found\"}\n")
// 				r.EXPECT().BalanceInfoUser(ctx, id).Return(uid, balance, err)
// 			},
// 			mockBehaviorWritingOff: func(r *mock_service.MockWritingOff, id int64, amount float32) {},
// 			expectedStatusCode:     400,
// 			expectedResponseBody:   "{\"error\":\"User not found\"}\n",
// 		},
// 	}
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			c := gomock.NewController(t)
// 			defer c.Finish()

// 			repoInfo := mock_service.NewMockBalanceInfoUser(c)
// 			test.mockBehaviorBalanceInfo(repoInfo, test.inputUser)

// 			repoWritingOff := mock_service.NewMockWritingOffUser(c)
// 			test.mockBehaviorWritingOff(repoWritingOff, test.inputUser, test.inputAmount)

// 			services := &service.Service{
// 				BalanceInfo: repoInfo,
// 				WritingOff:  repoWritingOff,
// 			}
// 			handler := Handler{services}

// 			r := chi.NewRouter()
// 			r.Patch("/writing-off", handler.WritingOff)

// 			w := httptest.NewRecorder()
// 			req := httptest.NewRequest(http.MethodPatch, "/writing-off",
// 				bytes.NewBufferString(test.inputBody))

// 			r.ServeHTTP(w, req)

// 			assert.Equal(t, w.Code, test.expectedStatusCode)
// 			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
// 		})
// 	}
// }
