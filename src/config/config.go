package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/kelseyhightower/envconfig"
)

var (
	// Config is a package variable, which is populated during init() execution and shared to whole application
	Config Configuration

	// FilePath defines a path to JSON-config file
	FilePath = "contact_book_service_cfg.json"

	// APIVersions stores slice of supported versions
	APIVersions = []string{"v1"}

	// Connections stores slice of all connections
	Connections = []string{"Cassandra"}
)

type (
	// Version options
	Version struct {
		SolutionName    string `json:"SolutionName"    envconfig:"CBS_SOLUTION_NAME"`
		ServiceName     string `json:"ServiceName"     envconfig:"CBS_SERVICE_NAME"`
		ServiceProvider string `json:"ServiceProvider" envconfig:"CBS_SERVICE_PROVIDER"`
	}

	// Configuration options
	Configuration struct {
		DebugHeaders      bool     `json:"DebugHeaders"          envconfig:"CBS_DEBUG_HEADERS"`
		FilePath          string   `json:"FilePath"              envconfig:"CBS_FILE_PATH"`
		LogFile           string   `json:"LogFilePath"           envconfig:"CBS_LOG_FILE"`
		ListenURL         string   `json:"ListenURL"             envconfig:"CBS_LISTEN_URL"`
		URLPrefix         string   `json:"URLPrefix"             envconfig:"CBS_URL_PREFIX"`
		APIVersion        string   `json:"APIVersion"            envconfig:"CBS_API_VERSION"`
		LogLevel          string   `json:"LogLevel"              envconfig:"CBS_LOG_LEVEL"`
		LogMaxFileSize    int64    `json:"LogMaxFileSize"        envconfig:"CBS_LOG_MAX_FILE_SIZE"`
		LogKeepOldFiles   int      `json:"LogKeepOldFiles"       envconfig:"CBS_LOG_KEEP_OLD_FILES"`
		CassandraHosts    []string `json:"CassandraHosts"        envconfig:"CBS_CASSANDRA_HOSTS"`
		CassandraKeyspace string   `json:"CassandraKeyspace"     envconfig:"CBS_CASSANDRA_KEYSPACE"`
		Version           Version  `json:"Version"`
		RequestTimeout    int      `json:"RequestTimeout"        envconfig:"CBS_REQUEST_TIMEOUT"`
		DefaultLangCBge   string   `json:"DefaultLangCBge"       envconfig:"CBS_DEFAULT_LANGCBGE"`
	}
)

// Load reads and loads configuration to Config variable
func Load() {
	var err error

	confLen := len(FilePath)
	if confLen != 0 {
		err = readFromJSON(FilePath)
	}
	if err == nil {
		err = readFromENV()
	}
	if err != nil {
		panic(`Configuration not found. Please specify configuration`)
	}
}

// isMissing validates Configuration
func (c *Configuration) isMissing() bool {
	return c.ListenURL == "" || len(c.CassandraHosts) == 0
}

// readFromJSON reads config data from JSON-file
func readFromJSON(configFilePath string) error {
	log.Printf("Looking for JSON config file (%s)", configFilePath)

	contents, err := ioutil.ReadFile(configFilePath)
	if err == nil {
		reader := bytes.NewBuffer(contents)
		err = json.NewDecoder(reader).Decode(&Config)
	}
	if err != nil {
		log.Printf("Reading configuration from JSON (%s) failed: %s\n", configFilePath, err)
	} else {
		log.Printf("Configuration has been read from JSON (%s) successfully\n", configFilePath)
	}

	return err
}

// readFromENV reads data from environment variables
func readFromENV() (err error) {
	log.Println("Looking for ENV configuration")

	err = envconfig.Process("PS", &Config)

	if err == nil && Config.isMissing() {
		err = errors.New("Configuration is missing")
	} else {
		log.Println("ENV configuration has been read successfully")
	}

	return err
}
