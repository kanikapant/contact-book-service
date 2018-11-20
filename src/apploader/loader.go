package apploader

import (
	"github.com/gocql/gocql"
	"github.com/kanikapant/common-lib/src/logging"
	"github.com/kanikapant/contact-book-service/src/config"
	"github.com/kanikapant/contact-book-service/src/logger"
	"github.com/kanikapant/contact-book-service/src/persistency/cassandra"
)

// AppLoader defines the whole configuration of the Application.
// It incorporates partial configurations of components
type AppLoader struct {
	ConfigService    config.Configuration
	LoggerService    logging.Logger
	CassandraService *gocql.ClusterConfig
}

// App is the whole configuration of the Application
var App *AppLoader

// LoadApplicationServices loads all partial configurations of components
// and populates the AppLoaderService with the configuration data
func LoadApplicationServices() {
	config.Load()
	if err := logger.Load(); err != nil {
		logger.Log.Errorf("failed to load logger: %s", err)
	}
	cassandra.Load()

	// TODO: add translation as App service:
	App = &AppLoader{
		ConfigService:    config.Config,
		LoggerService:    logger.Log,
		CassandraService: cassandra.CassandraClient,
	}

	App.CassandraService.Keyspace = App.ConfigService.CassandraKeyspace
	gocql.Logger = logger.CassandraLogger
}
