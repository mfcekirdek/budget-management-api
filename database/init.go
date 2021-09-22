package database

import (
	"github.com/couchbase/gocb/v2"
	"gitlab.com/mfcekirdek/budget-management-api/config"
	"log"
	"time"
)

var cluster *gocb.Cluster

func Setup(config *config.CouchbaseConfig) *gocb.Cluster {
	var err error
	cluster, err = gocb.Connect(config.Host, gocb.ClusterOptions{
		Username: config.User,
		Password: config.Pass,
	})
	if err != nil {
		log.Fatal(err)
	}
	return cluster
}

func GetCollection(bucketName string) *gocb.Collection {
	bucket := cluster.Bucket(bucketName)
	const timeOut = 10 * time.Second

	err := bucket.WaitUntilReady(timeOut, nil)
	if err != nil {
		log.Fatal(err)
	}
	return bucket.DefaultCollection()
}
