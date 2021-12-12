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

	userIdString := string(mapUser[id])
	uid, err := validate.IdValidate(userIdString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sortBy := mapUser[sortBy]
	if sortBy != nilValue && sortBy != data && sortBy != amount {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Errorf("Sort value passed")
		json.JSONError(w, "Sort value passed")
		return
	}

	orderBy := mapUser[orderBy]
	if orderBy != nilValue && orderBy != desc {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Errorf("Incorrect parameter for sorting, can only be created_at or amount")
		json.JSONError(w, "Incorrect parameter for sorting, can only be created_at or amount")
		return
	}

	descriptionSlice, err := h.services.GetDescriptions.GetDescriptionsUsers(ctx, uid, sortBy, orderBy)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.WithError(err).Errorf(err.Error())
		return
	}

	for _, row := range descriptionSlice {
		err := json.JSONUGetAllUsers(w, row)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logrus.WithError(err).Errorf(err.Error())
			return
		}
	}

}
