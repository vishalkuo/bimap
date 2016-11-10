package bimap

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


const key = "key"
const value = "value"

func TestNewBiMap(t *testing.T) {
	actual := NewBiMap()
	expected := biMap{forward:make(map[string]string), inverse:make(map[string]string)}
	assert.Equal(t, expected, *actual, "They should be equal")
}

func TestBiMap_Insert(t *testing.T) {
	actual := NewBiMap()
	actual.Insert(key, value)

	fwdExpected := make(map[string]string)
	invExpected := make(map[string]string)
	fwdExpected[key] = value
	invExpected[value] = key
	expected := biMap{forward:fwdExpected, inverse:invExpected}

	assert.Equal(t, expected, *actual, "They should be equal")
}

func TestBiMap_Exists(t *testing.T) {
	actual := NewBiMap()

	actual.Insert(key, value)
	assert.False(t, actual.Exists("ARBITARY_KEY"), "Key should not exist")
	assert.True(t, actual.Exists(key), "Inserted key should exist")
}

func TestBiMap_InverseExists(t *testing.T) {
	actual := NewBiMap()

	actual.Insert(key, value)
	assert.False(t, actual.InverseExists("ARBITARY_VALUE"), "Value should not exist")
	assert.True(t, actual.InverseExists(value), "Inserted value should exist")
}

func TestBiMap_Get(t *testing.T) {
	actual := NewBiMap()

	actual.Insert(key, value)

	actualVal, ok := actual.Get(key)

	assert.True(t, ok, "It should return true")
	assert.Equal(t, value, actualVal,  "Value and returned val should be equal")

	actualVal, ok = actual.Get(value)

	assert.False(t, ok, "It should return false")
	assert.Empty(t, actualVal, "Actual val should be empty")
}

func TestBiMap_GetInverse(t *testing.T) {
	actual := NewBiMap()

	actual.Insert(key, value)

	actualKey, ok := actual.InverseGet(value)

	assert.True(t, ok, "It should return true")
	assert.Equal(t, key, actualKey,"Key and returned key should be equal")

	actualKey, ok = actual.Get(value)

	assert.False(t, ok, "It should return false")
	assert.Empty(t, actualKey, "Actual key should be empty")
}

func TestBiMap_Size(t *testing.T) {
	actual := NewBiMap()

	assert.Equal(t, 0, actual.Size(), "Length of empty bimap should be zero")

	actual.Insert(key, value)

	assert.Equal(t, 1, actual.Size(), "Length of bimap should be one")
}

func TestBiMap_Delete(t *testing.T) {
	actual := NewBiMap()
	dummyKey := "DummyKey"
	dummyVal := "DummyVal"
	actual.Insert(key, value)
	actual.Insert(dummyKey, dummyVal)

	assert.Equal(t, 2, actual.Size(), "Size of bimap should be two")

	actual.Delete(dummyKey)

	fwdExpected := make(map[string]string)
	invExpected := make(map[string]string)
	fwdExpected[key] = value
	invExpected[value] = key

	expected := biMap{forward:fwdExpected, inverse:invExpected}

	assert.Equal(t, 1, actual.Size(), "Size of bimap should be two")
	assert.Equal(t, expected, *actual, "They should be the same")

	actual.Delete(dummyKey)

	assert.Equal(t, 1, actual.Size(), "Size of bimap should be two")
	assert.Equal(t, expected, *actual, "They should be the same")
}

func TestBiMap_InverseDelete(t *testing.T) {
	actual := NewBiMap()
	dummyKey := "DummyKey"
	dummyVal := "DummyVal"
	actual.Insert(key, value)
	actual.Insert(dummyKey, dummyVal)

	assert.Equal(t, 2, actual.Size(), "Size of bimap should be two")

	actual.InverseDelete(dummyVal)

	fwdExpected := make(map[string]string)
	invExpected := make(map[string]string)
	fwdExpected[key] = value
	invExpected[value] = key

	expected := biMap{forward:fwdExpected, inverse:invExpected}

	assert.Equal(t, 1, actual.Size(), "Size of bimap should be two")
	assert.Equal(t, expected, *actual, "They should be the same")

	actual.InverseDelete(dummyVal)

	assert.Equal(t, 1, actual.Size(), "Size of bimap should be two")
	assert.Equal(t, expected, *actual, "They should be the same")
}




