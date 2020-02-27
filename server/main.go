package main

import (
	"amp-templates/server/controllers"
	"amp-templates/server/middleware"
	"amp-templates/server/services/configuration"
	"amp-templates/server/services/log"
	"amp-templates/server/services/mux"
	"amp-templates/server/services/template_cache"
	"context"
	"github.com/NYTimes/gziphandler"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func makeHTTPServer() *http.Server {
	return &http.Server{
		Addr:           configuration.Config.Address,
		Handler:        mux.Router,
		ReadTimeout:    time.Duration(configuration.Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(configuration.Config.WriteTimeout * int64(time.Second)),
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func main() {
	controllers.ConfigureTemplateFuncs()
	template_cache.Configure()

	mux.Router.Use(middleware.LoggerMiddleware)
	mux.Router.Use(gziphandler.GzipHandler)

	log.Warning(configuration.Config.Application + " Server listening at " + configuration.Config.Address)

	httpServer := makeHTTPServer()

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Error("Failed to start server.")
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	_ = httpServer.Shutdown(ctx)
	log.Warning("Shutting down server.")
}
