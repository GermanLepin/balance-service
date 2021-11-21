package handlers_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
)

var (
	user     = "postgres"
	password = "1234"
	port     = "54320"
	dialect  = "postgres"
	dsn      = "postgres://%s:%s@localhost:%s/%s?sslmode=disable"
	idleConn = 25
	maxConn  = 25
)

type Repository interface {
	Close()
	Drop() error
}

type UserModel struct {
	ID    string
	Name  string
	Email string
	Phone string
}

func TestMai(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err.Error())
	}

	dsn = fmt.Sprintf(dsn, user, password, port, db)
	if err = pool.Retry(func() error {
		_, err := NewRepository(dialect, dsn, idleConn, maxConn)
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err.Error())
	}

	// defer func() {
	// 	repo.Close()
	// }()

	// err = repo.Drop()
	// if err != nil {
	// 	panic(err)
	// }

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

type repository struct {
	db *sql.DB
}

// NewRepository will create a variable that represent the Repository struct
func NewRepository(dialect, dsn string, idleConn, maxConn int) (Repository, error) {
	db, err := sql.Open(dialect, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(idleConn)
	db.SetMaxOpenConns(maxConn)

	return &repository{db}, nil
}

// Close attaches the provider and close the connection
func (r *repository) Close() {
	r.db.Close()
}

func (r *repository) Drop() error {
	ctx := context.Background()

	query := "DROP TABLE IF EXISTS users"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

// func TestUpBalance(t *testing.T) {
// 	JSONparams := bytes.NewBuffer([]byte(
// 		`{
// 			"id":"1",
// 			"amount":"1000.55"
// 			}`))

// 	req := httptest.NewRequest("POST", "localhost:9000/up-balance", JSONparams)
// 	r := httptest.NewRecorder()
// 	UpBalance(r, req)

// 	if status := r.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	result := string(body)
// 	expected := "{\"user id\":1,\"top up an amount\":1000.55}\n"

// 	if result != expected {
// 		t.Errorf("handler body: got %v want %s",
// 			result, expected)
// 	}
// }

// func TestUpBalanceErrorUserId(t *testing.T) {
// 	JSONparams := bytes.NewBuffer([]byte(
// 		`{
// 			"id":"-1",
// 			"amount":"10590.55"
// 		}`))

// 	req := httptest.NewRequest("POST", "localhost:9000/up-balance", JSONparams)
// 	r := httptest.NewRecorder()
// 	UpBalance(r, req)

// 	if status := r.Code; status != http.StatusBadRequest {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
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

// func TestUpBalanceErrorAmount(t *testing.T) {
// 	JSONparams := bytes.NewBuffer([]byte(
// 		`{
// 			"id":"1",
// 			"amount":"-10590.55"
// 		}`))

// 	req := httptest.NewRequest("POST", "localhost:9000/up-balance", JSONparams)
// 	r := httptest.NewRecorder()
// 	UpBalance(r, req)

// 	if status := r.Code; status != http.StatusBadRequest {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
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
