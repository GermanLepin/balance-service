package handler

import (
	"context"
	"net/http"
	"tech_task/pkg/helpers/parse"

	json "tech_task/pkg/helpers/json_responce"

	"github.com/sirupsen/logrus"
)

func (h *Handler) GetAllUsersDescriptionsSort(w http.ResponseWriter, r *http.Request) {
	mapUser, _ := parse.ParsJSON(r)
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

	if orderBy == nilValue && sortBy == nilValue {
		err := h.selectDataFromDatabase(ctx, w, data, asc, sqlOrderBy)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.JSONError(w, "User not found")
			return
		}
	}

	if orderBy == desc {
		switch {
		case sortBy == data:
			err := h.selectDataFromDatabase(ctx, w, data, desc, sqlOrderBy)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.JSONError(w, "User not found")
				return
			}

		case sortBy == amount:
			err := h.selectDataFromDatabase(ctx, w, amount, desc, sqlOrderBy)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.JSONError(w, "User not found")
				return
			}
		}
	}

	if orderBy == nilValue {
		switch {
		case sortBy == data:
			err := h.selectDataFromDatabase(ctx, w, data, asc, sqlOrderBy)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.JSONError(w, "User not found")
				return
			}

		case sortBy == amount:
			err := h.selectDataFromDatabase(ctx, w, amount, asc, sqlOrderBy)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.JSONError(w, "User not found")
				return
			}
		}
	}
}

func (h *Handler) selectDataFromDatabase(ctx context.Context, w http.ResponseWriter, sortBy, orderBy, sqlOrderBy string) error {
	descriptionSlice, err := h.services.GetAllUsersDescriptionsSort.GetAllDescriptionsSort(ctx, sortBy, orderBy, sqlOrderBy)
	if err != nil {
		return err
	}

	for _, row := range descriptionSlice {
		err := json.JSONUGetAllUsers(w, row)
		if err != nil {
			return err
		}
	}

	return nil
}
