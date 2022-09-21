package cache

import (
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/senpathi/kafkajet/internal/domain"
)

func TestClusterCache(t *testing.T) {
	clsCache := NewClusterCache()
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			cls := domain.Cluster{
				Name:    strconv.Itoa(i),
				Brokers: nil,
			}

			err := clsCache.Write(cls.Name, cls)
			assert.Equal(t, nil, err)

			v, ok := clsCache.Read(cls.Name)
			assert.Equal(t, true, ok)
			assert.Equal(t, cls, v)

			clsCache.Delete(cls.Name)
			_, ok = clsCache.Read(cls.Name)
			assert.Equal(t, false, ok)
		}(i)
	}

	wg.Wait()
}
