package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vision8tech/goings/common"
	"github.com/vision8tech/goings/shared/models"
	"log"
	"net/http"
	"strings"
)

//
// GetProjectsAPIHandler returns the list of available projects.
//
func GetProjectsAPIHandler(env *common.Env) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		projects, err := env.ProjectsRepo.RetrieveProjects()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		log.Printf("GetProjectsAPIHandler > result: %d entries", len(projects))
		_ = json.NewEncoder(w).Encode(projects)
	})

}

//
// GetProjectByIdAPIHandler returns the list of available projects.
//
func GetProjectByIdAPIHandler(env *common.Env) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		project, err := env.ProjectsRepo.RetrieveProjectByID(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if project == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(project)
	})

}

//
// SubmitProjectAPIHandler receives a POST request to create a new project.
//
func SubmitProjectAPIHandler(env *common.Env) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var project models.Project
		err := json.NewDecoder(r.Body).Decode(&project)
		defer func() { _ = r.Body.Close() }()
		if err != nil {
			log.Printf("SubmitProjectAPIHandler > Error decoding request body: %s", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		log.Printf("SubmitProjectAPIHandler > project=%#v\n", project)
		err = env.ProjectsRepo.StoreProject(&project)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			errText := err.Error()
			if strings.Contains(errText, "non-existent") {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			_ = json.NewEncoder(w).Encode(common.NewApiError(errText))
		}
	})

}

//
// UpdateProjectAPIHandler receives a POST request to create a new project.
//
func UpdateProjectAPIHandler(env *common.Env) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		var project models.Project
		err := json.NewDecoder(r.Body).Decode(&project)
		defer func() { _ = r.Body.Close() }()
		if err != nil {
			log.Printf("UpdateProjectAPIHandler > Error decoding request body: %s", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = env.ProjectsRepo.UpdateProjectByID(id, &project)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			errText := err.Error()
			if strings.Contains(errText, "non-existing") {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			_ = json.NewEncoder(w).Encode(common.NewApiError(errText))
		}
	})

}

//
// DeleteProjectAPIHandler receives a POST request to create a new project.
//
func DeleteProjectAPIHandler(env *common.Env) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		err := env.ProjectsRepo.DeleteProjectByID(id)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			errText := err.Error()
			if strings.Contains(errText, "non-existing") {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			_ = json.NewEncoder(w).Encode(common.NewApiError(errText))
		}
	})

}
