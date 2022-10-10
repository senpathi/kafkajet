package service

import (
	"github.com/senpathi/kafkajet/internal/datastore"
	"github.com/senpathi/kafkajet/internal/domain"
	"github.com/senpathi/kafkajet/internal/kafka"
	"sync"
)

type ClusterService struct {
	clusterRepo datastore.ClusterRepo
	clients     map[string]kafka.Client
	mu          *sync.Mutex
}

func NewService(clusterRepo datastore.ClusterRepo) *ClusterService {
	return &ClusterService{
		clusterRepo: clusterRepo,
		clients:     make(map[string]kafka.Client),
		mu:          new(sync.Mutex),
	}
}

func (s *ClusterService) ViewClusterTopics(clusterName string) ([]string, error) {
	cli, err := s.getClient(clusterName)
	if err != nil {
		return nil, err
	}

	return cli.Topics()
}

func (s *ClusterService) CreateTopics(clusterName string, details []domain.TopicDetails) ([]string, error) {
	cli, err := s.getClient(clusterName)
	if err != nil {
		return nil, err
	}

	return cli.CreateTopics(details)
}

func (s *ClusterService) getClient(name string) (kafka.Client, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	cli, ok := s.clients[name]
	if ok {
		return cli, nil
	}

	cluster, err := s.clusterRepo.Read(name)
	if err != nil {
		return nil, err
	}

	cli = kafka.NewClient(cluster)
	s.clients[name] = cli

	return cli, nil
}
