package goleveldb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGoLevelDB(t *testing.T) {
	db, err := NewDB(t.TempDir())
	require.NoError(t, err)

	// set key-value
	require.NoError(t, db.Put([]byte("key"), []byte("value")))
	require.NoError(t, db.Put([]byte("key1"), []byte("value1")))
	require.NoError(t, db.Put([]byte("key2"), []byte("value2")))
	require.NoError(t, db.Put([]byte("key3"), []byte("value3")))
	require.NoError(t, db.Put([]byte("xxxx"), []byte("value3")))

	// get value
	value, err := db.Get([]byte("key"))
	require.NoError(t, err)
	require.Equal(t, "value", string(value))

	// get next key-value
	key, value, err := db.Next([]byte("key0"))
	require.NoError(t, err)
	require.Equal(t, "key1", string(key))
	require.Equal(t, "value1", string(value))

	// get previous key-value
	key, value, err = db.Prev([]byte("key2"))
	require.NoError(t, err)
	require.Equal(t, "key1", string(key))
	require.Equal(t, "value1", string(value))
	key, value, err = db.Prev([]byte("key4"))
	require.NoError(t, err)
	require.Equal(t, "key3", string(key))
	require.Equal(t, "value3", string(value))

	// iterate all key-value
	count := 0
	require.NoError(t, db.Iterate([]byte("key"), func(key, value []byte) error {
		count++
		return nil
	}))
	require.Equal(t, 4, count)

	// prune
	require.NoError(t, db.Prune([]byte("key2")))
	_, err = db.Get([]byte("key1"))
	require.NoError(t, err)
	_, err = db.Get([]byte("key"))
	require.Error(t, err)

	// close db
	require.NoError(t, db.Close())
}
