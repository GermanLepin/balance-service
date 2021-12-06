package handler_test

import "testing"

func TestWritingOff(t *testing.T) {
	// 	JSONparams := bytes.NewBuffer([]byte(
	// 		`{
	// 			"id":"1",
	// 			"amount":"500.55"
	// 			}`))

	// 	req := httptest.NewRequest("POST", "localhost:9000/writing-off", JSONparams)
	// 	ctx := context.Background()
	// 	w := httptest.NewRecorder()
	// 	r := new(mockRepository)
	// 	r.On("BalanceInfoDB", ctx, w, 1).Return(1, 100.55)
	// 	// service := handler.HttpService{
	// 	// 	UserService: r,
	// 	// }
	// 	// service.WritingOff(w, req)
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
	// 	expected := "{\"user id\":1,\"writing off an amount\":500.55}\n"

	// 	if result != expected {
	// 		t.Errorf("handler body: got %v want %s",
	// 			result, expected)
	// 	}
	// }

	// func TestWritingOffErrorUserId(t *testing.T) {
	// 	JSONparams := bytes.NewBuffer([]byte(
	// 		`{
	// 			"id":"-1",
	// 			"amount":"10590.55"
	// 		}`))

	// 	req := httptest.NewRequest("POST", "localhost:9000/writing-off", JSONparams)
	// 	ctx := context.Background()
	// 	w := httptest.NewRecorder()
	// 	r := new(mockRepository)
	// 	r.On("BalanceInfoDB", ctx, w, 1).Return(1, 100.55)
	// 	service := handler.HttpService{
	// 		UserService: r,
	// 	}
	// 	service.WritingOff(w, req)

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

	// func TestWritingOffErrorAmount(t *testing.T) {
	// 	JSONparams := bytes.NewBuffer([]byte(
	// 		`{
	// 			"id":"1",
	// 			"amount":"-10590.55"
	// 		}`))

	// 	req := httptest.NewRequest("POST", "localhost:9000/writing-off", JSONparams)
	// 	ctx := context.Background()
	// 	w := httptest.NewRecorder()
	// 	r := new(mockRepository)
	// 	r.On("BalanceInfoDB", ctx, w, 1).Return(1, 100.55)
	// 	service := handler.HttpService{
	// 		UserService: r,
	// 	}
	// 	service.WritingOff(w, req)

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

	// func TestWritingOffErrorFindUserIdDB(t *testing.T) {
	// 	JSONparams := bytes.NewBuffer([]byte(
	// 		`{
	// 			"id":"1111111",
	// 			"amount":"10590.55"
	// 		}`))

	// 	req := httptest.NewRequest("POST", "localhost:9000/writing-off", JSONparams)
	// 	ctx := context.Background()
	// 	w := httptest.NewRecorder()
	// 	r := new(mockRepository)
	// 	r.On("BalanceInfoDB", ctx, w, 1).Return(1, 100.55)
	// 	service := handler.HttpService{
	// 		UserService: r,
	// 	}
	// 	service.WritingOff(w, req)

	// 	if status := w.Code; status != http.StatusBadRequest {
	// 		t.Errorf("handler returned wrong status code: got %v want %v",
	// 			status, http.StatusOK)
	// 	}

	// 	body, err := ioutil.ReadAll(w.Body)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	result := string(body)
	// 	expected := "{\"error\":\"User not found in database\"}\n"

	// 	if result != expected {
	// 		t.Errorf("handler body: got %v want %v",
	// 			result, expected)
	// 	}
}
