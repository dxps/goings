package api

import (
	"encoding/json"
	"github.com/vision8tech/goings/shared/models"
	"log"
	"net/http"
	"time"

	"github.com/vision8tech/goings/common"
)

// GetProjectsAPIEndpoint returns the handler for "/api/projects" endpoint.
func GetProjectsAPIEndpoint(env *common.Env) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		projects := env.ProjectsRepo.RetrieveProjects()
		w.Header().Set("Content-Type", "application/json")
		log.Printf("GetProjectsAPIEndpoint > result: %d entries", len(projects))
		json.NewEncoder(w).Encode(projects)
	})

}

//
func PostProjectAPIEndpoint(env *common.Env) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		project := models.Project{ID: "test1", Title: "Project Test 1", Description: "desc test",
			ImageURI: "http://some.url", StartTime: time.Now(), State: "0",
		}
		log.Printf("PostProjectAPIEndpoint > project=%v\n", project)
		w.Header().Set("Content-Type", "application/json")
		env.ProjectsRepo.StoreProject(&project)
	})

}
