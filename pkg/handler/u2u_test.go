package handler

// import (
// 	"bytes"
// 	"context"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestU2U(t *testing.T) {
// 	JSONparams := bytes.NewBuffer([]byte(
// 		`{
// 			"id1":"1",
// 			"id2":"2",
// 			"amount":"60"
// 			}`))

// 	req := httptest.NewRequest("POST", "localhost:9000/user-to-user", JSONparams)
// 	ctx := context.Background()
// 	w := httptest.NewRecorder()
// 	r := new(mockRepository)
// 	r.On("BalanceInfoDB", ctx, w, 1).Return(1, 100.55)
// 	// service := handler.HttpService{
// 	// 	UserService: r,
// 	// }

// 	//service.U2U(w, req)
// 	fmt.Println(req)

// 	if status := w.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	body, err := ioutil.ReadAll(w.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	result := string(body)
// 	expected := "{\"user id sender\":1,\"writing off an amount\":60,\"user id recipient\":2}\n"

// 	if result != expected {
// 		t.Errorf("handler body: got %v want %s",
// 			result, expected)
// 	}
// }

// func TestU2UErrorUserId(t *testing.T) {
// 	JSONparams := bytes.NewBuffer([]byte(
// 		`{
// 			"id1":"1",
// 			"id2":"-2",
// 			"amount":"10590.55"
// 		}`))

// 	req := httptest.NewRequest("POST", "localhost:9000/user-to-user", JSONparams)
// 	ctx := context.Background()
// 	w := httptest.NewRecorder()
// 	r := new(mockRepository)
// 	r.On("BalanceInfoDB", ctx, w, 1).Return(1, 100.55)
// 	// service := handler.HttpService{
// 	// 	UserService: r,
// 	// }

// 	// service.U2U(w, req)
// 	fmt.Println(req)

// 	if status := w.Code; status != http.StatusBadRequest {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	body, err := ioutil.ReadAll(w.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	result := string(body)
// 	expected := "{\"error\":\"Incorrect value id user\"}\n"

// 	if result != expected {
// 		t.Errorf("handler body: got %v want %v",
// 			result, expected)
// 	}
// }

// func TestU2UErrorAmount(t *testing.T) {
// 	JSONparams := bytes.NewBuffer([]byte(
// 		`{
// 			"id1":"1",
// 			"id2":"2",
// 			"amount":"-10590.55"
// 		}`))

// 	req := httptest.NewRequest("POST", "localhost:9000/user-to-user", JSONparams)
// 	ctx := context.Background()
// 	w := httptest.NewRecorder()
// 	r := new(mockRepository)
// 	r.On("BalanceInfoDB", ctx, w, 1).Return(1, 100.55)
// 	// service := handler.HttpService{
// 	// 	UserService: r,
// 	// }

// 	// service.U2U(w, req)
// 	fmt.Println(req)

// 	if status := w.Code; status != http.StatusBadRequest {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	body, err := ioutil.ReadAll(w.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	result := string(body)
// 	expected := "{\"error\":\"The amount is negative\"}\n"

// 	if result != expected {
// 		t.Errorf("handler body: got %v want %v",
// 			result, expected)
// 	}
// }

// func TestU2UErrorFindUserIdDB(t *testing.T) {
// 	JSONparams := bytes.NewBuffer([]byte(
// 		`{
// 			"id1":"1111111111",
// 			"id2":"1",
// 			"amount":"10590.55"
// 		}`))

// 	req := httptest.NewRequest("POST", "localhost:9000/user-to-user", JSONparams)
// 	ctx := context.Background()
// 	w := httptest.NewRecorder()
// 	r := new(mockRepository)
// 	r.On("BalanceInfoDB", ctx, w, 1).Return(1, 100.55)
// 	// service := handler.HttpService{
// 	// 	UserService: r,
// 	// }
// 	// service.U2U(w, req)

// 	fmt.Println(req)
// 	if status := w.Code; status != http.StatusBadRequest {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	body, err := ioutil.ReadAll(w.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	result := string(body)
// 	expected := "{\"error\":\"User not found\"}\n"

// 	if result != expected {
// 		t.Errorf("handler body: got %v want %v",
// 			result, expected)
// 	}
// }

// package handler

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

// func TestHandler_WritingOff(t *testing.T) {
// 	type mockBehaviorInfo func(r *mock_service.MockBalanceInfo, id int64)
// 	type mockBehaviorWritingOff func(r *mock_service.MockWritingOff, id int64, amount float64)
// 	ctx := context.Background()

// 	tests := []struct {
// 		name                   string
// 		inputBody              string
// 		inputUser              int64
// 		inputAmount            float64
// 		mockBehaviorInfo       mockBehaviorInfo
// 		mockBehaviorWritingOff mockBehaviorWritingOff
// 		expectedStatusCode     int
// 		expectedResponseBody   string
// 	}{
// 		{
// 			name:        "Ok",
// 			inputBody:   `{"id":"1","amount":"969.63"}`,
// 			inputUser:   1,
// 			inputAmount: 969.63,
// 			mockBehaviorInfo: func(r *mock_service.MockBalanceInfo, id int64) {
// 				var uid int64 = 1
// 				var balance float64 = 1830.55
// 				var err error = nil
// 				r.EXPECT().BalanceInfoUser(ctx, id).Return(uid, balance, err)
// 			},
// 			mockBehaviorWritingOff: func(r *mock_service.MockWritingOff, id int64, amount float64) {
// 				var uid int64 = 1
// 				var respAmount float64 = 969.63
// 				var err error = nil
// 				r.EXPECT().WritingOffUser(ctx, id, amount).Return(uid, respAmount, err)
// 			},
// 			expectedStatusCode:   200,
// 			expectedResponseBody: "{\"user id\":1,\"writing off an amount\":969.63}\n",
// 		},
// 		// {
// 		// 	name:                 "Wrong input user",
// 		// 	inputBody:            `{"id":"-1","amount":"1569.77"}`,
// 		// 	inputUser:            -1,
// 		// 	inputAmount:          1569.77,
// 		// 	mockBehavior:         func(r *mock_service.MockWritingOff, id int64, amount float64) {},
// 		// 	expectedStatusCode:   400,
// 		// 	expectedResponseBody: "{\"error\":\"incorrect value id user\"}\n",
// 		// },
// 		// {
// 		// 	name:                 "Wrong input amount",
// 		// 	inputBody:            `{"id":"1","amount":"-1569.77"}`,
// 		// 	inputUser:            1,
// 		// 	inputAmount:          -1569.77,
// 		// 	mockBehavior:         func(r *mock_service.MockWritingOff, id int64, amount float64) {},
// 		// 	expectedStatusCode:   400,
// 		// 	expectedResponseBody: "{\"error\":\"the amount is negative\"}\n",
// 		// },
// 		// {
// 		// 	name:                 "Wrong input more 2 decimal places",
// 		// 	inputBody:            `{"id":"1","amount":"1569.77345"}`,
// 		// 	inputUser:            1,
// 		// 	inputAmount:          1569.77345,
// 		// 	mockBehavior:         func(r *mock_service.MockWritingOff, id int64, amount float64) {},
// 		// 	expectedStatusCode:   400,
// 		// 	expectedResponseBody: "{\"error\":\"the amount have more then 2 decimal places\"}\n",
// 		// },
// 		// {
// 		// 	name:                 "User not found",
// 		// 	inputBody:            `{"id":"1","amount":"1569.77"}`,
// 		// 	inputUser:            99999999,
// 		// 	inputAmount:          1569.77,
// 		// 	mockBehavior:         func(r *mock_service.MockWritingOff, id int64, amount float64) {},
// 		// 	expectedStatusCode:   400,
// 		// 	expectedResponseBody: "{\"error\":\"user not found\"}\n",
// 		// },
// 	}
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			c := gomock.NewController(t)
// 			defer c.Finish()

// 			repoInfo := mock_service.NewMockBalanceInfoUser(c)
// 			test.mockBehaviorInfo(repoInfo, test.inputUser)

// 			repoWritingOff := mock_service.NewMockWritingOffUser(c)
// 			test.mockBehaviorWritingOff(repoWritingOff, test.inputUser, test.inputAmount)

// 			servicesInfo := &service.Service{BalanceInfo: repoInfo}
// 			handler := Handler{servicesInfo}

// 			servicesWritingOff := &service.Service{WritingOff: repoWritingOff}
// 			handler = Handler{servicesWritingOff}

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
