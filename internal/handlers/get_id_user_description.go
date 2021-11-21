package handlers

import (
	"net/http"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/parseform"
	"tech_task/pkg/helpers/validate"

	log "github.com/sirupsen/logrus"
)

func GetUserId(w http.ResponseWriter, r *http.Request) {
	mapUser := parseform.ParsJSON(r)
	userIdString := string(mapUser[id])
	userId := validate.IdValidate(w, userIdString)
	if userId < 1 {
		return
	}

	descriptionSlice := instance.GetUserIdDB(ctx, w, userId)
	if descriptionSlice == nil {
		return
	}

	for _, row := range descriptionSlice {
		jsonenc.JSONUGetUserIdDescription(w, row)
	}
}

func GetUserIdSort(w http.ResponseWriter, r *http.Request) {
	mapUser := parseform.ParsJSON(r)
	userIdString := string(mapUser[id])
	userId := validate.IdValidate(w, userIdString)
	if userId < 1 {
		return
	}

	sortBy := string(mapUser[sortBy])
	if sortBy == nilValue {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Sort value passed")
		jsonenc.JSONError(w, "Sort value passed")
		return
	}

	if sortBy != amount && sortBy != data && sortBy != nilValue {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Incorrect parameter for sorting, can only be created_at or amount")
		jsonenc.JSONError(w, "Incorrect parameter for sorting, can only be created_at or amount")
		return
	}

	switch sortBy {
	case data:
		descriptionSlice := instance.GetUserIdSortDB(ctx, w, userId, data)
		if descriptionSlice == nil {
			return
		}

		for _, row := range descriptionSlice {
			jsonenc.JSONUGetUserIdDescription(w, row)
		}

	case amount:
		descriptionSlice := instance.GetUserIdSortDB(ctx, w, userId, amount)
		if descriptionSlice == nil {
			return
		}

		for _, row := range descriptionSlice {
			jsonenc.JSONUGetUserIdDescription(w, row)
		}
	}
}

func GetUserIdSortDesc(w http.ResponseWriter, r *http.Request) {
	mapUser := parseform.ParsJSON(r)
	userIdString := string(mapUser[id])
	userId := validate.IdValidate(w, userIdString)
	if userId < 1 {
		return
	}

	sortBy := string(mapUser[sortBy])
	if sortByDesc == nilValue {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Sort value passed")
		jsonenc.JSONError(w, "Sort value passed")
		return
	}

	if sortBy != amount && sortBy != data && sortBy != nilValue {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Incorrect parameter for sorting, can only be created_at or amount")
		jsonenc.JSONError(w, "Incorrect parameter for sorting, can only be created_at or amount")
		return
	}

	switch sortByDesc {
	case data:
		descriptionSlice := instance.GetUserIdSortDescDB(ctx, w, userId, data)
		if descriptionSlice == nil {
			return
		}

		for _, row := range descriptionSlice {
			jsonenc.JSONUGetUserIdDescription(w, row)
		}

	case amount:
		descriptionSlice := instance.GetUserIdSortDescDB(ctx, w, userId, amount)
		if descriptionSlice == nil {
			return
		}

		for _, row := range descriptionSlice {
			jsonenc.JSONUGetUserIdDescription(w, row)
		}
	}
}
