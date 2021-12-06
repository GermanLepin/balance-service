package handler

import (
	"net/http"
	"tech_task/pkg/helpers/parse"

	json "tech_task/pkg/helpers/json_responce"
)

func (h *Handler) GetAllUsersDescriptionsSort(w http.ResponseWriter, r *http.Request) {
	mapUser, err := parse.ParsJSON(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sortBy := mapUser[sortBy]
	orderBy := mapUser[orderBy]

	switch {
	case mapUser[sortBy] == "":
		descriptionSlice, err := h.services.GetAllUsersDescriptionsSort.GetAllDescriptionsSort(ctx, amount, asc)
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
	case sortBy == data:
		descriptionSlice, err := h.services.GetAllUsersDescriptionsSort.GetAllDescriptionsSort(ctx, data, asc)
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

	case sortBy == amount:
		descriptionSlice, err := h.services.GetAllUsersDescriptionsSort.GetAllDescriptionsSort(ctx, amount, asc)
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
	case sortBy == data && orderBy == desc:
		descriptionSlice, err := h.services.GetAllUsersDescriptionsSort.GetAllDescriptionsSort(ctx, data, desc)
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
	case sortBy == amount && orderBy == desc:
		descriptionSlice, err := h.services.GetAllUsersDescriptionsSort.GetAllDescriptionsSort(ctx, amount, desc)
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
