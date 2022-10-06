package cache

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/senpathi/kafkajet/internal/domain"
	domainErr "github.com/senpathi/kafkajet/internal/errors"
)

func NewClusterCache() Cache {
	return &clusterCache{
		mu:       new(sync.RWMutex),
		clusters: make(map[string]domain.Cluster),
	}
}

// clusterCache caches cluster details
type clusterCache struct {
	mu       *sync.RWMutex
	clusters map[string]domain.Cluster
}

func (c *clusterCache) Read(id string) (value interface{}, exists bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cluster, ok := c.clusters[id]

	return cluster, ok
}

func (c *clusterCache) Write(id string, value interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	cluster, ok := value.(domain.Cluster)
	if !ok {
		return domainErr.Error{
			Message: "invalid value type",
			Err:     fmt.Errorf("invalid type, expected [domain.Cluster], received [%v]", reflect.TypeOf(value)),
			Code:    domainErr.InvalidValueTypeErrorCode,
		}
	}

	c.clusters[id] = cluster

	return nil
}

func (c *clusterCache) Delete(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.clusters, id)
}
