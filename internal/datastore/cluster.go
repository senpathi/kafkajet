package datastore

import (
	"github.com/senpathi/kafkajet/internal/datastore/cache"
	"github.com/senpathi/kafkajet/internal/datastore/repository"
	"github.com/senpathi/kafkajet/internal/domain"
)

type ClusterRepo interface {
	Read(name string) (value domain.Cluster, err error)
	Write(cluster domain.Cluster) error
	Delete(name string) error
}

func NewClusterRepository() ClusterRepo {
	c := cache.NewClusterCache()
	c.Write("localhost", domain.Cluster{
		Name:    "localhost",
		Brokers: []string{"kafka:9092"},
	})

	return &clusterRepository{
		cache: c,
		table: database.Table(`cluster`),
	}
}

type clusterRepository struct {
	cache cache.Cache
	table repository.Repo
}

func (c *clusterRepository) Read(name string) (value domain.Cluster, err error) {
	cluster, ok := c.cache.Read(name)
	if ok {
		return cluster.(domain.Cluster), nil
	}

	return domain.Cluster{}, repository.ErrorNotFound
}

func (c *clusterRepository) Write(cluster domain.Cluster) error {
	//TODO implement me
	panic("implement me")
}

func (c *clusterRepository) Delete(name string) error {
	//TODO implement me
	panic("implement me")
}
