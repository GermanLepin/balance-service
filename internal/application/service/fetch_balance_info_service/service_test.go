package fetch_balance_info_service

// import (
// 	"bytes"
// 	"context"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/go-chi/chi"
// 	"github.com/golang/mock/gomock"
// 	"gotest.tools/assert"
// )

// func TestHandler_BalanceInfo(t *testing.T) {
// 	type mockBehavior func(r *mock_service.MockBalanceInfo, id int64)
// 	ctx := context.Background()

// 	tests := []struct {
// 		name                 string
// 		inputBody            string
// 		inputUser            int64
// 		mockBehavior         mockBehavior
// 		expectedStatusCode   int
// 		expectedResponseBody string
// 	}{
// 		{
// 			name:      "Ok",
// 			inputBody: `{"user id":"1"}`,
// 			inputUser: 1,
// 			mockBehavior: func(r *mock_service.MockBalanceInfo, id int64) {
// 				var uid int64 = 1
// 				var balance = 830.55
// 				var err error = nil
// 				r.EXPECT().BalanceInfoUser(ctx, id).Return(uid, balance, err)
// 			},
// 			expectedStatusCode:   200,
// 			expectedResponseBody: "{\"user id\":1,\"balance\":830.55}\n",
// 		},
// 		{
// 			name:                 "Wrong Input",
// 			inputBody:            `{"user id":"-1"}`,
// 			inputUser:            -1,
// 			mockBehavior:         func(r *mock_service.MockBalanceInfo, id int64) {},
// 			expectedStatusCode:   400,
// 			expectedResponseBody: "{\"error\":\"incorrect value user id\"}\n",
// 		},
// 		{
// 			name:      "User not found",
// 			inputBody: `{"user id":"99999999"}`,
// 			inputUser: 99999999,
// 			mockBehavior: func(r *mock_service.MockBalanceInfo, id int64) {
// 				var uid int64 = 0
// 				var balance float32 = 0
// 				var err = errors.New("{\"error\":\"User not found\"}\n")
// 				r.EXPECT().BalanceInfoUser(ctx, id).Return(uid, balance, err)
// 			},
// 			expectedStatusCode:   400,
// 			expectedResponseBody: "{\"error\":\"User not found\"}\n",
// 		},
// 	}
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			c := gomock.NewController(t)
// 			defer c.Finish()

// 			repo := mock_service.NewMockBalanceInfoUser(c)
// 			test.mockBehavior(repo, test.inputUser)

// 			services := &service.Service{BalanceInfo: repo}
// 			handler := Handler{services}

// 			r := chi.NewRouter()
// 			r.Get("/balance-info", handler.BalanceInfo)

// 			w := httptest.NewRecorder()
// 			req := httptest.NewRequest(http.MethodGet, "/balance-info",
// 				bytes.NewBufferString(test.inputBody))

// 			r.ServeHTTP(w, req)

// 			assert.Equal(t, w.Code, test.expectedStatusCode)
// 			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
// 		})
// 	}
// }
