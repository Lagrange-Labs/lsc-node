package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	maxCount := uint64(10)
	c := NewCache(maxCount)

	upTo := uint64(20)
	for i := uint64(0); i < upTo; i++ {
		if i >= maxCount {
			_, ok := c.Get(i - maxCount)
			require.True(t, ok, "key %d should be in the cache", i-maxCount)
		}
		c.Set(i, i)
	}

	for i := uint64(0); i < upTo-maxCount*2; i++ {
		_, ok := c.Get(i)
		require.False(t, ok, "key %d should not be in the cache", i)
	}

	for i := upTo - maxCount*2; i < upTo; i++ {
		ret, ok := c.Get(i)
		require.True(t, ok, "key %d should be in the cache", i)
		require.Equal(t, i, ret)
	}
}
