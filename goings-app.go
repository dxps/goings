package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/vision8tech/goings/common"
	"github.com/vision8tech/goings/repos"
	//"github.com/vision8tech/goings/common"
)

var appRoot string
var appServerPort string
var staticAssetsPath string

func main() {

	env := common.Env{}
	var projectsRepo repos.ProjectsRepo = repos.NewSqliteProjectRepo()

	log.Printf("%d exists.", projectsRepo.GetProjects())

	router := mux.NewRouter()

	registerRoutes(&env, router)

	// Register Request Handler for Static Assetcs
	fs := http.FileServer(http.Dir(staticAssetsPath))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.Handle("/", router)

	http.ListenAndServe(":"+appServerPort, nil)

}

func init() {

	appRoot = os.Getenv("GOINGS_APP_ROOT")
	staticAssetsPath = appRoot + "/static"

}

// registerRoutes is responsible for regisetering the server-side request handlers
func registerRoutes(env *common.Env, r *mux.Router) {

	// todo

}
