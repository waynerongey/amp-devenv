package controllers

import (
	"amp-templates/server/services/log"
	"amp-templates/server/services/mux"
	"net/http"
)

const (
	HomeRoute string = "home"
)


func init() {
	ConfigureTemplateFuncs()
	ConfigureRoutes()
}

//ConfigureRoutes configures our routes..duh.
func ConfigureRoutes() {
	//404 Not Found Route
	mux.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Warning(r.RemoteAddr, " Tried to get path: ", r.URL.Path)
		http.Redirect(w, r, RebuildRoute(HomeRoute), http.StatusSeeOther)
	})

	//Home
	mux.Router.HandleFunc("/", Home).Methods("Get").Name(HomeRoute)
	mux.Router.HandleFunc("/api/dev/hr", HotReload).Methods("Get")
}
