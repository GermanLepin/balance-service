package handler

import (
	"net/http"
	json "tech_task/pkg/helpers/json_responce"
	"tech_task/pkg/helpers/parse"
	"tech_task/pkg/helpers/validate"

	"github.com/sirupsen/logrus"
)

func (h *Handler) UpBalance(w http.ResponseWriter, r *http.Request) {
	mapUser, err := parse.ParsJSON(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userIdString := string(mapUser[id])
	amountString := string(mapUser[amount])

	userId, err := validate.IdValidate(userIdString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.JSONError(w, err.Error())
		return
	}

	correctAmount, err := validate.AmountValidate(amountString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.JSONError(w, err.Error())
		return
	}

	err = h.services.UpBalance.UpBalanceUser(ctx, userId, correctAmount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.WithError(err).Errorf("user not found")
		return
	}

	err = json.JSONUpBalance(w, userId, correctAmount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
