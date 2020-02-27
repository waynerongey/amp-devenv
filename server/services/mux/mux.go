package mux

import (
	"amp-templates/server/services/configuration"
	"github.com/gorilla/mux"
	"net/http"
)

//Router is the multiplexing router from gorilla
var Router *mux.Router

func init() {
	Router = mux.NewRouter()

	//Static Files
	Router.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir(configuration.Config.Static))))

	http.Handle("/", Router)
}
