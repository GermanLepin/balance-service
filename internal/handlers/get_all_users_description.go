package handlers

import (
	"net/http"
	"tech_task/pkg/helpers/jsonenc"
	"tech_task/pkg/helpers/parseform"
	"tech_task/pkg/helpers/pg"

	log "github.com/sirupsen/logrus"
)

var (
	instance = pg.StartDB()
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	descriptionSlice := instance.GetAllUsersDB(ctx, w)
	if descriptionSlice == nil {
		return
	}

	for _, row := range descriptionSlice {
		jsonenc.JSONUGetAllUsers(w, row)
	}
}

func GetAllUsersSort(w http.ResponseWriter, r *http.Request) {
	sortBy := parseform.Pars(w, r, sortBy)
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

func GetAllUsersSortDesc(w http.ResponseWriter, r *http.Request) {
	sortByDesc := parseform.Pars(w, r, sortByDesc)
	if sortByDesc != amount && sortByDesc != data && sortByDesc != nilValue {
		w.WriteHeader(http.StatusBadRequest)
		log.Errorf("Incorrect parameter for sorting, can only be created_at or amount")
		jsonenc.JSONError(w, "Incorrect parameter for sorting, can only be created_at or amount")
		return
	}

	switch sortByDesc {
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
