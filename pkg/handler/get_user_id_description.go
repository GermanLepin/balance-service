package handler

import (
	"context"
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
		err := h.selectDataFromDatabaseUserId(ctx, w, userId, data, asc, sqlOrderBy)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.JSONError(w, "User not found")
			return
		}
	}

	if orderBy == desc {
		switch {
		case sortBy == data:
			err := h.selectDataFromDatabaseUserId(ctx, w, userId, data, desc, sqlOrderBy)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.JSONError(w, "User not found")
				return
			}

		case sortBy == amount:
			err := h.selectDataFromDatabaseUserId(ctx, w, userId, amount, desc, sqlOrderBy)
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
			err := h.selectDataFromDatabaseUserId(ctx, w, userId, data, asc, sqlOrderBy)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.JSONError(w, "User not found")
				return
			}

		case sortBy == amount:
			err := h.selectDataFromDatabaseUserId(ctx, w, userId, amount, asc, sqlOrderBy)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.JSONError(w, "User not found")
				return
			}
		}
	}
}

func (h *Handler) selectDataFromDatabaseUserId(ctx context.Context, w http.ResponseWriter, userId int64, sortBy, orderBy, sqlOrderBy string) error {
	descriptionSlice, err := h.services.GetUserIdDescriptionsSort.GetUserIdDescriptionsSort(ctx, userId, sortBy, orderBy, sqlOrderBy)
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
