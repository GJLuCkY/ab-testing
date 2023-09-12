package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type response struct {
	ExperimentID int    `json:"experiment_id"`
	GroupID      int    `json:"group_id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	UserID       string `json:"user_id"`
}

func (s service) GetGroupByUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		slug := vars["slug"]
		userID := r.URL.Query().Get("user_id")

		experiment, err := s.ExperimentService.GetBySlug(r.Context(), slug)
		if err != nil {
			s.respond(w, err, 0)
			return
		}

		experimentGroups, err := s.ExperimentService.GetActiveGroups(r.Context(), experiment.ID)
		if err != nil {
			s.respond(w, err, 0)
			return
		}

		if len(experimentGroups) == 0 {
			s.respond(w, map[string]string{"error_message": "У эксперимента нет групп, или они выключены"}, 422)
			return
		}

		experimentGroup, err := s.ExperimentService.GetGroup(r.Context(), experiment.ID, userID)
		if err != nil {
			s.respond(w, err, 0)
			return
		}

		s.respond(w, response{
			ExperimentID: experimentGroup.ExperimentId,
			GroupID:      experimentGroup.ID,
			Name:         experimentGroup.Name,
			Slug:         experimentGroup.Slug,
			UserID:       userID,
		}, http.StatusOK)
	}
}
