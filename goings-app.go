package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"go.isomorphicgo.org/go/isokit"

	"github.com/gorilla/mux"
	"github.com/vision8tech/goings/api/endpoints"
	"github.com/vision8tech/goings/common"
	"github.com/vision8tech/goings/handlers"
	"github.com/vision8tech/goings/repos"
	"github.com/vision8tech/goings/shared/templatefuncs"
)

var appMode string
var appRoot string
var appServerPort string
var staticAssetsPath string

func init() {

	appMode = os.Getenv("GOINGS_APP_MODE")
	appRoot = os.Getenv("GOINGS_APP_ROOT")
	appServerPort = os.Getenv("GOINGS_APP_PORT")
	staticAssetsPath = appRoot + "/static"
	log.Printf("[main] init > Env vars: GOINGS_APP_ROOT=%s GOINGS_APP_PORT=%s", appRoot, appServerPort)
}

// Initialization of the repositories.
func initRepos(env *common.Env) {

	var projectsRepo repos.ProjectsRepo = repos.NewSqliteProjectRepo()
	env.ProjectsRepo = projectsRepo

}

// Initialization of the templates used in server-side rendering.
func initTemplateSet(env *common.Env, generateStaticAssets bool) {

	isokit.WebAppRoot = appRoot
	isokit.StaticAssetsPath = staticAssetsPath
	isokit.StaticTemplateBundleFilePath = staticAssetsPath + "/templates/app.tmplbundle"
	isokit.TemplateFilesPath = appRoot + "/shared/templates"
	isokit.TemplateFileExtension = ".gohtml"
	isokit.PrefixNamePartial = "parts/"

	templateSet := isokit.NewTemplateSet()
	templateSet.Funcs = template.FuncMap{
		"rubyformat":     templatefuncs.RubyDate,
		"unixformat":     templatefuncs.UnixTime,
		"productionmode": templatefuncs.IsProduction}
	templateSet.GatherTemplates()
	env.TemplateSet = templateSet

}

func main() {

	log.SetFlags(log.Ldate | log.Lmicroseconds)

	env := common.Env{}
	initRepos(&env)
	initTemplateSet(&env, false)

	router := mux.NewRouter()
	registerRoutes(&env, router)

	// Register Request Handler for Static Assetcs
	fs := http.FileServer(http.Dir(staticAssetsPath))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.Handle("/", router)

	http.ListenAndServe(":"+appServerPort, nil)

}

// registerRoutes is responsible for registering the server-side request handlers
func registerRoutes(env *common.Env, r *mux.Router) {

	// Client-side triggered request handlers.
	if appMode != "production" {
		r.Handle("/js/ui.js", handlers.GopherjsScriptHandlerExt(appRoot, "/ui/ui.js")).Methods("GET")
		r.Handle("/js/ui.js.map", handlers.GopherjsScriptMapHandlerExt(appRoot, "/ui/ui.js.map")).Methods("GET")
	}

	// Register handler for the delivery of the template bundle.
	r.Handle("/template-bundle", handlers.TemplateBundleHandler(env)).Methods("POST")

	// REST API Endpoints
	r.Handle("/api/projects", endpoints.GetProjectsAPIEndpoint(env)).Methods("GET")

	// Back-end Pages
	r.Handle("/", handlers.IndexHandler(env)).Methods("GET")

	log.Println("[main] registerRoutes > Routes registered.")

}
