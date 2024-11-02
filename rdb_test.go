package rdb_test

import (
	"os"
	"testing"

	rdb "github.com/ross96D/go_rdb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatabase(t *testing.T) {
	db, err := rdb.New("test_db")
	require.NoError(t, err)
	defer os.Remove("test_db")

	db.Insert("key1", []byte("value1"))
	db.Insert("key2", []byte("value2"))
	db.Insert("key3", []byte("value3"))
	db.Insert("key4", []byte("value4"))

	assert.Equal(t, []byte("value1"), db.Search("key1"))
	assert.Equal(t, []byte("value2"), db.Search("key2"))
	assert.Equal(t, []byte("value3"), db.Search("key3"))
	assert.Equal(t, []byte("value4"), db.Search("key4"))

	db.Update("key1", []byte("val1"))
	db.Update("key2", []byte("val2"))
	db.Update("key3", []byte("val3"))
	db.Update("key4", []byte("val4"))

	assert.Equal(t, []byte("val1"), db.Search("key1"))
	assert.Equal(t, []byte("val2"), db.Search("key2"))
	assert.Equal(t, []byte("val3"), db.Search("key3"))
	assert.Equal(t, []byte("val4"), db.Search("key4"))

	db.Delete("key1")
	db.Delete("key2")
	db.Delete("key3")
	db.Delete("key4")

	assert.Equal(t, []byte(nil), db.Search("key1"))
	assert.Equal(t, []byte(nil), db.Search("key2"))
	assert.Equal(t, []byte(nil), db.Search("key3"))
	assert.Equal(t, []byte(nil), db.Search("key4"))
}

func TestDatabaseASD(t *testing.T) {
	_, err := rdb.New("not/a/valid/path")
	require.Error(t, err)
}
