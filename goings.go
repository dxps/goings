package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"go.isomorphicgo.org/go/isokit"

	"github.com/gorilla/mux"

	"github.com/vision8tech/goings/common"
	"github.com/vision8tech/goings/handlers/api"
	"github.com/vision8tech/goings/handlers/pages"
	"github.com/vision8tech/goings/handlers/ui"
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
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Printf("main.init > Env vars: GOINGS_APP_ROOT=%s GOINGS_APP_PORT=%s", appRoot, appServerPort)

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

	env := common.Env{}
	initRepos(&env)
	initTemplateSet(&env, false)

	router := mux.NewRouter()
	registerRoutes(&env, router)

	http.Handle("/", router)
	http.ListenAndServe(":"+appServerPort, nil)

}

// registerRoutes is responsible for registering the server-side request handlers
func registerRoutes(env *common.Env, r *mux.Router) {

	// Standard/Initial requests handlers (for pages, not views)

	r.Handle("/", pages.IndexPageHandler(env)).Methods("GET")

	// UI (client-side) triggered request handlers

	r.Handle("/template-bundle", ui.TemplateBundleHandler(env)).Methods("POST")

	// API request handlers

	r.Handle("/api/projects", api.GetProjectsAPIEndpoint(env)).Methods("GET")
	r.Handle("/api/projects/{id}", api.GetProjectByIdAPIEndpoint(env)).Methods("GET")
	r.Handle("/api/projects", api.SubmitProjectAPIEndpoint(env)).Methods("POST")

	// ----------------------------------------------------------------------

	// static assets requests handlers
	fs := http.FileServer(http.Dir(staticAssetsPath))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	log.Println("main.registerRoutes > Routes registered.")

}
