package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/kanikapant/contact-book-service/src/apploader"
	"github.com/kanikapant/contact-book-service/src/config"
	"github.com/kanikapant/contact-book-service/src/models"
	"github.com/kanikapant/contact-book-service/src/services"
	"github.com/urfave/negroni"
)

func main() {

	configFile := flag.String("config", "contact_book_service_cfg.json", "Configuration file in JSON-format")
	flag.Parse()

	config.FilePath = *configFile

	//recovery log in case user account service crash
	defer func() {
		if r := recover(); r != nil {
			apploader.App.LoggerService.Errorf("contact book service crashed. Error Details:- %s  Call Stack:- %s", r, fmt.Sprintf("%s", debug.Stack()))
		}
	}()

	// initialising global Application Service
	apploader.LoadApplicationServices()

	// TODO: move into cassandra.Load() take into account /health endpoint
	cassandraConfig := models.GetCassandraConfig()
	apploader.App.LoggerService.Debugf("Cassandra config:- %s ", cassandraConfig)

	// setting up web server middlewares
	middlewareManager := negroni.New()
	middlewareManager.Use(negroni.NewRecovery())
	middlewareManager.UseHandler(services.NewRouter())

	apploader.App.LoggerService.Info("Service started")

	err := http.ListenAndServe(apploader.App.ConfigService.ListenURL, middlewareManager)
	apploader.App.LoggerService.Info("Stop running application: %s", err)

}
