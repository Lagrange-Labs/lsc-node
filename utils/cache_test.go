package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	maxCount := uint64(10)
	c := NewCache(maxCount)

	upTo := uint64(20)
	for i := uint64(0); i < upTo; i++ {
		c.Set(i, i)
	}

	for i := uint64(0); i < upTo-maxCount; i++ {
		_, ok := c.Get(i)
		require.False(t, ok, "key %d should not be in the cache", i)
	}

	for i := upTo - maxCount; i < upTo; i++ {
		ret, ok := c.Get(i)
		require.True(t, ok, "key %d should be in the cache", i)
		require.Equal(t, i, ret)
	}
}
