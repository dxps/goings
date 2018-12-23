package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vision8tech/goings/common"
	"github.com/vision8tech/goings/shared/models"
	"log"
	"net/http"
)

// GetProjectsAPIEndpoint returns the list of available projects.
func GetProjectsAPIEndpoint(env *common.Env) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		projects, err := env.ProjectsRepo.RetrieveProjects()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		log.Printf("GetProjectsAPIEndpoint > result: %d entries", len(projects))
		_ = json.NewEncoder(w).Encode(projects)
	})

}

// GetProjectsAPIEndpoint returns the list of available projects.
func GetProjectByIdAPIEndpoint(env *common.Env) http.Handler {

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

// SubmitProjectAPIEndpoint receives a POST request to create a new project.
func SubmitProjectAPIEndpoint(env *common.Env) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var project models.Project
		err := json.NewDecoder(r.Body).Decode(&project)
		defer func() { _ = r.Body.Close() }()
		if err != nil {
			log.Printf("SubmitProjectAPIEndpoint > Error decoding request body: %s", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		log.Printf("SubmitProjectAPIEndpoint > project=%#v\n", project)
		env.ProjectsRepo.StoreProject(&project)
	})

}
