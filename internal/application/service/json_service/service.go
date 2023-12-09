package json_service

import (
	"balance-service/internal/application/dto"
	"encoding/json"
	"log"
	"net/http"
)

func (s *service) ErrorJSON(w http.ResponseWriter, err error, statusCode int) error {
	var payload dto.JsonResponse
	payload.Message = err.Error()

	log.Printf("error: %s\n", err)

	out, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

type service struct{}

func New() *service {
	return &service{}
}
