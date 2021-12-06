package handler

import (
	"net/http"
	json "tech_task/pkg/helpers/json_responce"
	"tech_task/pkg/helpers/parse"
	"tech_task/pkg/helpers/validate"

	"github.com/sirupsen/logrus"
)

func (h *Handler) GetUserIdDescriptionsSort(w http.ResponseWriter, r *http.Request) {
	mapUser, err := parse.ParsJSON(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userIdString := string(mapUser[id])
	userId, err := validate.IdValidate(userIdString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.JSONError(w, err.Error())
		return
	}

	sortBy := string(mapUser[sortBy])
	if sortBy == nilValue {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Errorf("Sort value passed")
		json.JSONError(w, "Sort value passed")
		return
	}

	if sortBy != amount && sortBy != data && sortBy != nilValue {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Errorf("Incorrect parameter for sorting, can only be created_at or amount")
		json.JSONError(w, "Incorrect parameter for sorting, can only be created_at or amount")
		return
	}

	switch sortBy {
	case data:
		descriptionSlice, err := h.services.GetUserIdDescriptionsSort.GetUserIdDescriptionsSort(ctx, userId, data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		for _, row := range descriptionSlice {
			err := json.JSONUGetAllUsers(w, row)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

	case amount:
		descriptionSlice, err := h.services.GetUserIdDescriptionsSort.GetUserIdDescriptionsSort(ctx, userId, amount)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		for _, row := range descriptionSlice {
			err := json.JSONUGetAllUsers(w, row)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
