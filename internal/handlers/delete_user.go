package handlers

import (
	"net/http"
	"tech_task/pkg/helpers/jsonenc.go"
	"tech_task/pkg/helpers/pg"
	"tech_task/pkg/helpers/validate"

	log "github.com/sirupsen/logrus"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	instance := pg.StartDB()

	id := validate.IdValidate(w, r, id)
	if id < 1 {
		return
	}

	userId, balance := instance.BalanceInfoDB(ctx, w, id)
	if userId == 0 {
		return
	}

	if balance != 0.00 {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("The balance is not equal 0")
		jsonenc.JSONError(w, "The balance is not equal 0")
		return
	}

	instance.DeleteUserDB(ctx, id)
	jsonenc.JSONDeleteUser(w, id, "User deleted successfully")
}
