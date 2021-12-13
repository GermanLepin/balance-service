package handler

import (
	"net/http"
	"tech_task/pkg/helpers/parse"
	"tech_task/pkg/helpers/validate"

	json "tech_task/pkg/helpers/json_responce"

	"github.com/sirupsen/logrus"
)

func (h *Handler) GetDescriptions(w http.ResponseWriter, r *http.Request) {
	mapUser, _ := parse.ParsJSON(r)
	userIDString := string(mapUser[id])

	var uid int64 = 0
	if userIDString != "" {
		userID, err := validate.IdValidate(userIDString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.JSONError(w, err.Error())
			return
		}
		uid = userID
	}

	sortBy := mapUser[sortBy]
	if sortBy != nilValue && sortBy != data && sortBy != amount {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Errorf("Incorrect parameter for sorting, can only be created_at or amount")
		json.JSONError(w, "Incorrect parameter for sorting, can only be created_at or amount")
		return
	}

	orderBy := mapUser[orderBy]
	if orderBy != nilValue && orderBy != desc && orderBy != asc {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Errorf("Incorrect parameter for order, can only be desc or asc")
		json.JSONError(w, "Incorrect parameter for order, can only be desc or asc")
		return
	}

	descriptionSlice, err := h.services.GetDescriptions.GetDescriptionsUsers(ctx, uid, sortBy, orderBy)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.WithError(err).Errorf(err.Error())
		json.JSONError(w, err.Error())
		return
	}

	for _, row := range descriptionSlice {
		err := json.JSONGetDescriptions(w, row)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logrus.WithError(err).Errorf(err.Error())
			return
		}
	}
}
