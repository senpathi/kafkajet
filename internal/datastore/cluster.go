package datastore

import (
	"github.com/senpathi/kafkajet/internal/datastore/cache"
	"github.com/senpathi/kafkajet/internal/datastore/repository"
)

func NewClusterRepository() repository.Repo {
	return clusterRepository{
		cache: cache.NewClusterCache(),
		table: database.Table(`cluster`),
	}
}

type clusterRepository struct {
	cache cache.Cache
	table repository.Repo
}

func (c clusterRepository) Read(filter map[string]interface{}) (value interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (c clusterRepository) Write(value interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c clusterRepository) Delete(filter map[string]interface{}) {
	//TODO implement me
	panic("implement me")
}
