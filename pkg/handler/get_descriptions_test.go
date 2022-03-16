package handler

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"tech_task"
	"tech_task/pkg/service"
	mock_service "tech_task/pkg/service/mocks"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestHandler_GetAllUsersDescriptionsSort(t *testing.T) {
	type mockBehaviorGetDescriptions func(r *mock_service.MockGetDescriptions, id int64, sortBy, orderBy string)
	ctx := context.Background()

	tests := []struct {
		name                        string
		inputBody                   string
		uid                         int64
		sortBy                      string
		orderBy                     string
		mockBehaviorGetDescriptions mockBehaviorGetDescriptions
		expectedStatusCode          int
		expectedResponseBody        string
	}{
		{
			name:      "Get all descriptions",
			inputBody: "",
			uid:       0,
			sortBy:    "",
			orderBy:   "",
			mockBehaviorGetDescriptions: func(r *mock_service.MockGetDescriptions, id int64, sortBy, orderBy string) {
				var list = []tech_task.Description{
					{
						Id:              1,
						SenderReceiver:  "Avito",
						Amount:          6780,
						Description:     "Покупка наушников",
						BalanceAtMoment: 1520,
						UserID:          1,
						CreatedAt:       time.Date(2021, 10, 28, 01, 17, 8, 660784, time.UTC),
						Refill:          "F",
					},
					{
						Id:              2,
						SenderReceiver:  "Avito",
						Amount:          5490,
						Description:     "Продажа куртки",
						BalanceAtMoment: 7010,
						UserID:          1,
						CreatedAt:       time.Date(2021, 10, 28, 01, 32, 37, 122076, time.UTC),
						Refill:          "T",
					},
					{
						Id:              4,
						SenderReceiver:  "Avito",
						Amount:          7490,
						Description:     "Продажа зеркала",
						BalanceAtMoment: 9190,
						UserID:          2,
						CreatedAt:       time.Date(2021, 10, 28, 01, 32, 56, 434473, time.UTC),
						Refill:          "T",
					},
				}
				var err error = nil
				r.EXPECT().GetDescriptionsUsers(ctx, id, sortBy, orderBy).Return(list, err)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":1,\"sender receiver\":\"Avito\",\"amount\":6780,\"description\":\"Покупка наушников\",\"balance at moment\":1520,\"user id\":1,\"created at\":\"2021-10-28T01:17:08.000660784Z\",\"refill\":\"F\"}\n{\"id\":2,\"sender receiver\":\"Avito\",\"amount\":5490,\"description\":\"Продажа куртки\",\"balance at moment\":7010,\"user id\":1,\"created at\":\"2021-10-28T01:32:37.000122076Z\",\"refill\":\"T\"}\n{\"id\":4,\"sender receiver\":\"Avito\",\"amount\":7490,\"description\":\"Продажа зеркала\",\"balance at moment\":9190,\"user id\":2,\"created at\":\"2021-10-28T01:32:56.000434473Z\",\"refill\":\"T\"}\n",
		},
		{
			name:      "Get user descriptions",
			inputBody: `{"user id":"1"}`,
			uid:       1,
			sortBy:    "",
			orderBy:   "",
			mockBehaviorGetDescriptions: func(r *mock_service.MockGetDescriptions, id int64, sortBy, orderBy string) {
				var list = []tech_task.Description{
					{
						Id:              1,
						SenderReceiver:  "Avito",
						Amount:          6780,
						Description:     "Покупка наушников",
						BalanceAtMoment: 1520,
						UserID:          1,
						CreatedAt:       time.Date(2021, 10, 28, 01, 17, 8, 660784, time.UTC),
						Refill:          "F",
					},
					{
						Id:              2,
						SenderReceiver:  "Avito",
						Amount:          5490,
						Description:     "Продажа куртки",
						BalanceAtMoment: 7010,
						UserID:          1,
						CreatedAt:       time.Date(2021, 10, 28, 01, 32, 37, 122076, time.UTC),
						Refill:          "T",
					},
				}
				var err error = nil
				r.EXPECT().GetDescriptionsUsers(ctx, id, sortBy, orderBy).Return(list, err)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":1,\"sender receiver\":\"Avito\",\"amount\":6780,\"description\":\"Покупка наушников\",\"balance at moment\":1520,\"user id\":1,\"created at\":\"2021-10-28T01:17:08.000660784Z\",\"refill\":\"F\"}\n{\"id\":2,\"sender receiver\":\"Avito\",\"amount\":5490,\"description\":\"Продажа куртки\",\"balance at moment\":7010,\"user id\":1,\"created at\":\"2021-10-28T01:32:37.000122076Z\",\"refill\":\"T\"}\n",
		},
		{
			name:      "Get user descriptions sort amount",
			inputBody: `{"user id":"1","sort by":"amount"}`,
			uid:       1,
			sortBy:    "amount",
			orderBy:   "",
			mockBehaviorGetDescriptions: func(r *mock_service.MockGetDescriptions, id int64, sortBy, orderBy string) {
				var list = []tech_task.Description{
					{
						Id:              2,
						SenderReceiver:  "Avito",
						Amount:          5490,
						Description:     "Продажа куртки",
						BalanceAtMoment: 7010,
						UserID:          1,
						CreatedAt:       time.Date(2021, 10, 28, 01, 32, 37, 122076, time.UTC),
						Refill:          "T",
					},
					{
						Id:              1,
						SenderReceiver:  "Avito",
						Amount:          6780,
						Description:     "Покупка наушников",
						BalanceAtMoment: 1520,
						UserID:          1,
						CreatedAt:       time.Date(2021, 10, 28, 01, 17, 8, 660784, time.UTC),
						Refill:          "F",
					},
				}
				var err error = nil
				r.EXPECT().GetDescriptionsUsers(ctx, id, sortBy, orderBy).Return(list, err)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":2,\"sender receiver\":\"Avito\",\"amount\":5490,\"description\":\"Продажа куртки\",\"balance at moment\":7010,\"user id\":1,\"created at\":\"2021-10-28T01:32:37.000122076Z\",\"refill\":\"T\"}\n{\"id\":1,\"sender receiver\":\"Avito\",\"amount\":6780,\"description\":\"Покупка наушников\",\"balance at moment\":1520,\"user id\":1,\"created at\":\"2021-10-28T01:17:08.000660784Z\",\"refill\":\"F\"}\n",
		},
		{
			name:      "Get user descriptions sort amount order by desc",
			inputBody: `{"user id":"1","sort by":"amount", "order by":"desc"}`,
			uid:       1,
			sortBy:    "amount",
			orderBy:   "desc",
			mockBehaviorGetDescriptions: func(r *mock_service.MockGetDescriptions, id int64, sortBy, orderBy string) {
				var list = []tech_task.Description{
					{
						Id:              1,
						SenderReceiver:  "Avito",
						Amount:          6780,
						Description:     "Покупка наушников",
						BalanceAtMoment: 1520,
						UserID:          1,
						CreatedAt:       time.Date(2021, 10, 28, 01, 17, 8, 660784, time.UTC),
						Refill:          "F",
					},
					{
						Id:              2,
						SenderReceiver:  "Avito",
						Amount:          5490,
						Description:     "Продажа куртки",
						BalanceAtMoment: 7010,
						UserID:          1,
						CreatedAt:       time.Date(2021, 10, 28, 01, 32, 37, 122076, time.UTC),
						Refill:          "T",
					},
				}
				var err error = nil
				r.EXPECT().GetDescriptionsUsers(ctx, id, sortBy, orderBy).Return(list, err)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":1,\"sender receiver\":\"Avito\",\"amount\":6780,\"description\":\"Покупка наушников\",\"balance at moment\":1520,\"user id\":1,\"created at\":\"2021-10-28T01:17:08.000660784Z\",\"refill\":\"F\"}\n{\"id\":2,\"sender receiver\":\"Avito\",\"amount\":5490,\"description\":\"Продажа куртки\",\"balance at moment\":7010,\"user id\":1,\"created at\":\"2021-10-28T01:32:37.000122076Z\",\"refill\":\"T\"}\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMocketGetDescriptions(c)
			test.mockBehaviorGetDescriptions(repo, test.uid, test.sortBy, test.orderBy)

			services := &service.Service{
				GetDescriptions: repo,
			}
			handler := Handler{services}

			r := chi.NewRouter()
			r.Get("/descriptions/get", handler.GetDescriptions)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/descriptions/get",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
