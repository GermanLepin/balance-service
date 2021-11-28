package handler

import (
	"net/http"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/parseform"

	log "github.com/sirupsen/logrus"
)

func (h *HttpService) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	descriptionSlice := instance.GetAllUsersDB(ctx, w)
	if descriptionSlice == nil {
		return
	}

	for _, row := range descriptionSlice {
		jsonenc.JSONUGetAllUsers(w, row)
	}
}

func (h *HttpService) GetAllUsersSort(w http.ResponseWriter, r *http.Request) {
	mapUser := parseform.ParsJSON(r)
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

	switch sortBy {
	case data:
		descriptionSlice := instance.GetAllUsersSortDB(ctx, w, data)
		if descriptionSlice == nil {
			return
		}

		for _, row := range descriptionSlice {
			jsonenc.JSONUGetAllUsers(w, row)
		}

	case amount:
		descriptionSlice := instance.GetAllUsersSortDB(ctx, w, amount)
		if descriptionSlice == nil {
			return
		}

		for _, row := range descriptionSlice {
			jsonenc.JSONUGetAllUsers(w, row)
		}
	}
}

func (h *HttpService) GetAllUsersSortDesc(w http.ResponseWriter, r *http.Request) {
	mapUser := parseform.ParsJSON(r)
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

	switch sortBy {
	case data:
		descriptionSlice := instance.GetAllUsersSortDescDB(ctx, w, data)
		if descriptionSlice == nil {
			return
		}

		for _, row := range descriptionSlice {
			jsonenc.JSONUGetAllUsers(w, row)
		}

	case amount:
		descriptionSlice := instance.GetAllUsersSortDescDB(ctx, w, amount)
		if descriptionSlice == nil {
			return
		}

		for _, row := range descriptionSlice {
			jsonenc.JSONUGetAllUsers(w, row)
		}
	}
}
