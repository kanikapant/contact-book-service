package cassandra

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/kanikapant/contact-book-service/src/config"
)

// CassandraClient is a component providing access to Cassandra storage
var CassandraClient *gocql.ClusterConfig

// Load creates a Cassandra Client and populates it with initial data
func Load() {
	clusterConfig := gocql.NewCluster(config.Config.CassandraHosts...)
	clusterConfig.ProtoVersion = 4
	clusterConfig.Timeout = 3 * time.Second
	CassandraClient = clusterConfig
	CassandraClient.Consistency = gocql.Quorum
}
