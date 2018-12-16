package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/vision8tech/goings/api/endpoints"
	"github.com/vision8tech/goings/common"
	"github.com/vision8tech/goings/repos"
)

var appRoot string
var appServerPort string
var staticAssetsPath string

func init() {

	appRoot = os.Getenv("GOINGS_APP_ROOT")
	appServerPort = os.Getenv("GOINGS_APP_PORT")
	staticAssetsPath = appRoot + "/static"
}

func main() {

	var projectsRepo repos.ProjectsRepo = repos.NewSqliteProjectRepo()
	log.Printf("%d exists.", projectsRepo.GetProjects())

	env := common.Env{}
	env.ProjectsRepo = projectsRepo

	router := mux.NewRouter()
	registerRoutes(&env, router)
	log.Println("Routes registered.")

	// Register Request Handler for Static Assetcs
	fs := http.FileServer(http.Dir(staticAssetsPath))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.Handle("/", router)

	http.ListenAndServe(":"+appServerPort, nil)

}

// registerRoutes is responsible for registering the server-side request handlers
func registerRoutes(env *common.Env, r *mux.Router) {

	// REST API Endpoints
	r.Handle("/api/projects", endpoints.GetProjectsAPIEndpoint(env)).Methods("GET")

	// TODO: routes to be registered

}
