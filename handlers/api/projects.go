package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vision8tech/goings/common"
)

// GetProjectsAPIEndpoint returns the handler for "/api/projects" endpoint.
func GetProjectsAPIEndpoint(env *common.Env) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		projects := env.ProjectsRepo.GetProjects()
		w.Header().Set("Content-Type", "application/json")
		log.Printf("[/api/projects] result: %d", projects)
		json.NewEncoder(w).Encode(projects)
	})

}
