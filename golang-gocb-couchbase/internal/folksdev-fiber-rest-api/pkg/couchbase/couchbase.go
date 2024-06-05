package couchbase

import (
	"fmt"
	"github.com/couchbase/gocb/v2"
	"time"
)

//ConnectCluster

func ConnectCluster(host, username, password string, connBufferSizeInBytes uint) *gocb.Cluster {
	cluster, err := gocb.Connect(
		host,
		gocb.ClusterOptions{
			Username: username,
			Password: password,
			TimeoutsConfig: gocb.TimeoutsConfig{
				KVTimeout: 10 * time.Second,
			},
			InternalConfig: gocb.InternalConfig{
				ConnectionBufferSize: connBufferSizeInBytes,
			},
		})

	if err != nil {
		panic(fmt.Sprintf("error connecting to couchbase cluster: %s", err.Error()))
	}

	return cluster
}
