package main

import (
	"github.com/gorilla/mux"
	"github.com/vision8tech/goings/common"
	"github.com/vision8tech/goings/handlers/api"
	"github.com/vision8tech/goings/handlers/pages"
	"github.com/vision8tech/goings/handlers/ui"
	"github.com/vision8tech/goings/repos/sqlite"
	"github.com/vision8tech/goings/shared/templatefuncs"
	"go.isomorphicgo.org/go/isokit"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// application settings
var appMode string
var appRoot string
var appServerPort string

// http server settings
var server http.Server
var staticAssetsPath string
var idleConnsClosed chan struct{}

//
// init is used for the global application initialization.
//
func init() {

	appMode = os.Getenv("GOINGS_APP_MODE")
	appRoot = os.Getenv("GOINGS_APP_ROOT")
	appServerPort = os.Getenv("GOINGS_APP_PORT")
	staticAssetsPath = appRoot + "/static"
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Printf("main.init > Env vars: GOINGS_APP_ROOT=%s GOINGS_APP_PORT=%s", appRoot, appServerPort)
	idleConnsClosed = make(chan struct{})

}

//
// initRepos is initializing the repositories.
//
func initRepos(env *common.Env) {

	env.ProjectsRepo = sqlite.NewSqliteProjectRepo()

}

//
// initTemplateSet is initializing the templates used in server-side rendering cases.
//
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

//
// MAIN
//
func main() {

	env := common.Env{}
	initRepos(&env)
	initTemplateSet(&env, false)

	router := mux.NewRouter()
	registerRoutes(&env, router)
	http.Handle("/", router)

	setupGracefulShutdownActions(&env)

	err := http.ListenAndServe(":"+appServerPort, nil)
	if err != nil {
		log.Printf("main > Error starting the http server: '%s'\n", err)
	}

}

//
// registerRoutes is responsible for registering the server-side request handlers
//
func registerRoutes(env *common.Env, r *mux.Router) {

	// Standard/Initial requests handlers (for pages, not views)

	r.Handle("/", pages.IndexPageHandler(env)).Methods("GET")

	// --------------------------------------------------------------------------------------------

	// UI (client-side) triggered request handlers

	r.Handle("/template-bundle", ui.TemplateBundleHandler(env)).Methods("POST")

	// --------------------------------------------------------------------------------------------

	// API request handlers

	r.Handle("/api/projects", api.GetProjectsAPIHandler(env)).Methods(http.MethodGet)
	r.Handle("/api/projects", api.SubmitProjectAPIHandler(env)).Methods(http.MethodPost)
	r.Handle("/api/projects/{id}", api.GetProjectByIdAPIHandler(env)).Methods(http.MethodGet)
	r.Handle("/api/projects/{id}", api.UpdateProjectAPIHandler(env)).Methods(http.MethodPut)
	r.Handle("/api/projects/{id}", api.DeleteProjectAPIHandler(env)).Methods(http.MethodDelete)

	// --------------------------------------------------------------------------------------------

	// static assets requests handlers
	fs := http.FileServer(http.Dir(staticAssetsPath))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	log.Println("main.registerRoutes > Routes ready.")

}

//
// setupGracefulShutdownActions is setting up the actions in case of a graceful shutdown.
//
func setupGracefulShutdownActions(env *common.Env) {

	var gracefulStopChan = make(chan os.Signal)
	signal.Notify(gracefulStopChan, syscall.SIGTERM)
	signal.Notify(gracefulStopChan, syscall.SIGINT)
	go func() {
		sig := <-gracefulStopChan
		log.Printf("\nmain > Shutting down ('%+v' signal received) ...\n", sig)
		env.ProjectsRepo.Uninit()
		os.Exit(0)
	}()

}
