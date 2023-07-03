package simplecache_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	simplecache "github.com/zakirkun/infra-go/pkg/cache/simple_cache"
)

func TestSimpleCache_SetAndGet(t *testing.T) {
	simpleCache := simplecache.NewSimpleCache(simplecache.SimpleCache{
		ExpiredAt: 10,
		PurgeTime: 30,
	})

	cache := simpleCache.Open()
	simpleCache.Set("hello", "world")

	value := simpleCache.Get("hello")
	assert.NotNil(t, value)
	assert.Equal(t, "world", *value)

	cache.Flush()
}

func TestSimpleCache_Delete(t *testing.T) {
	simpleCache := simplecache.NewSimpleCache(simplecache.SimpleCache{
		ExpiredAt: 10,
		PurgeTime: 30,
	})

	cache := simpleCache.Open()
	simpleCache.Set("hello", "world")

	simpleCache.Delete("hello")

	value := simpleCache.Get("world")
	assert.Nil(t, value)

	cache.Flush()
}
