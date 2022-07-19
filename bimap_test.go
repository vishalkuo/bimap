package bimap

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const key = "key"
const value = "value"

func TestNewBiMap(t *testing.T) {
	actual := NewBiMap[string, string]()
	expected := &BiMap[string, string]{forward: make(map[string]string), inverse: make(map[string]string)}
	assert.Equal(t, expected, actual, "They should be equal")
}

func TestNewBiMapFrom(t *testing.T) {
	actual := NewBiMapFromMap(map[string]string{
		key: value,
	})
	actual.Insert(key, value)

	fwdExpected := make(map[string]string)
	invExpected := make(map[string]string)
	fwdExpected[key] = value
	invExpected[value] = key
	expected := &BiMap[string, string]{forward: fwdExpected, inverse: invExpected}

	assert.Equal(t, expected, actual, "They should be equal")
}

func TestNewBiMapFrom(t *testing.T) {
	actual := NewBiMapFrom(map[interface{}]interface{}{
		key: value,
	})
	actual.Insert(key, value)

	fwdExpected := make(map[interface{}]interface{})
	invExpected := make(map[interface{}]interface{})
	fwdExpected[key] = value
	invExpected[value] = key
	expected := &BiMap{forward: fwdExpected, inverse: invExpected}

	assert.Equal(t, expected, actual, "They should be equal")
}

func TestBiMap_Insert(t *testing.T) {
	actual := NewBiMap[string, string]()
	actual.Insert(key, value)

	fwdExpected := make(map[string]string)
	invExpected := make(map[string]string)
	fwdExpected[key] = value
	invExpected[value] = key
	expected := &BiMap[string, string]{forward: fwdExpected, inverse: invExpected}

	assert.Equal(t, expected, actual, "They should be equal")
}

func TestBiMap_Exists(t *testing.T) {
	actual := NewBiMap[string, string]()

	actual.Insert(key, value)
	assert.False(t, actual.Exists("ARBITARY_KEY"), "Key should not exist")
	assert.True(t, actual.Exists(key), "Inserted key should exist")
}

func TestBiMap_InverseExists(t *testing.T) {
	actual := NewBiMap[string, string]()

	actual.Insert(key, value)
	assert.False(t, actual.ExistsInverse("ARBITARY_VALUE"), "Value should not exist")
	assert.True(t, actual.ExistsInverse(value), "Inserted value should exist")
}

func TestBiMap_Get(t *testing.T) {
	actual := NewBiMap[string, string]()

	actual.Insert(key, value)

	actualVal, ok := actual.Get(key)

	assert.True(t, ok, "It should return true")
	assert.Equal(t, value, actualVal, "Value and returned val should be equal")

	actualVal, ok = actual.Get(value)

	assert.False(t, ok, "It should return false")
	assert.Empty(t, actualVal, "Actual val should be empty")
}

func TestBiMap_GetInverse(t *testing.T) {
	actual := NewBiMap[string, string]()

	actual.Insert(key, value)

	actualKey, ok := actual.GetInverse(value)

	assert.True(t, ok, "It should return true")
	assert.Equal(t, key, actualKey, "Key and returned key should be equal")

	actualKey, ok = actual.Get(value)

	assert.False(t, ok, "It should return false")
	assert.Empty(t, actualKey, "Actual key should be empty")
}

func TestBiMap_Size(t *testing.T) {
	actual := NewBiMap[string, string]()

	assert.Equal(t, 0, actual.Size(), "Length of empty bimap should be zero")

	actual.Insert(key, value)

	assert.Equal(t, 1, actual.Size(), "Length of bimap should be one")
}

func TestBiMap_Delete(t *testing.T) {
	actual := NewBiMap[string, string]()
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

	expected := &BiMap[string, string]{forward: fwdExpected, inverse: invExpected}

	assert.Equal(t, 1, actual.Size(), "Size of bimap should be two")
	assert.Equal(t, expected, actual, "They should be the same")

	actual.Delete(dummyKey)

	assert.Equal(t, 1, actual.Size(), "Size of bimap should be two")
	assert.Equal(t, expected, actual, "They should be the same")
}

func TestBiMap_InverseDelete(t *testing.T) {
	actual := NewBiMap[string, string]()
	dummyKey := "DummyKey"
	dummyVal := "DummyVal"
	actual.Insert(key, value)
	actual.Insert(dummyKey, dummyVal)

	assert.Equal(t, 2, actual.Size(), "Size of bimap should be two")

	actual.DeleteInverse(dummyVal)

	fwdExpected := make(map[string]string)
	invExpected := make(map[string]string)
	fwdExpected[key] = value
	invExpected[value] = key

	expected := &BiMap[string, string]{forward: fwdExpected, inverse: invExpected}

	assert.Equal(t, 1, actual.Size(), "Size of bimap should be two")
	assert.Equal(t, expected, actual, "They should be the same")

	actual.DeleteInverse(dummyVal)

	assert.Equal(t, 1, actual.Size(), "Size of bimap should be two")
	assert.Equal(t, expected, actual, "They should be the same")
}

func TestBiMap_WithVaryingType(t *testing.T) {
	actual := NewBiMap[string, int]()
	dummyKey := "Dummy key"
	dummyVal := 3

	actual.Insert(dummyKey, dummyVal)

	res, _ := actual.Get(dummyKey)
	resVal, _ := actual.GetInverse(dummyVal)
	assert.Equal(t, dummyVal, res, "Get by string key should return integer val")
	assert.Equal(t, dummyKey, resVal, "Get by integer val should return string key")

}

func TestBiMap_MakeImmutable(t *testing.T) {
	actual := NewBiMap[string, int]()
	dummyKey := "Dummy key"
	dummyVal := 3

	actual.Insert(dummyKey, dummyVal)

	actual.MakeImmutable()

	assert.Panics(t, func() {
		actual.Delete(dummyKey)
	}, "It should panic on a mutation operation")

	val, _ := actual.Get(dummyKey)

	assert.Equal(t, dummyVal, val, "It should still have the value")

	assert.Panics(t, func() {
		actual.DeleteInverse(dummyVal)
	}, "It should panic on a mutation operation")

	key, _ := actual.GetInverse(dummyVal)

	assert.Equal(t, dummyKey, key, "It should still have the key")

	size := actual.Size()

	assert.Equal(t, 1, size, "Size should be one")

	assert.Panics(t, func() {
		actual.Insert("New", 1)
	}, "It should panic on a mutation operation")

	size = actual.Size()

	assert.Equal(t, 1, size, "Size should be one")

}

func TestBiMap_GetForwardMap(t *testing.T) {
	actual := NewBiMap[string, int]()
	dummyKey := "Dummy key"
	dummyVal := 42

	forwardMap := make(map[string]int)
	forwardMap[dummyKey] = dummyVal

	actual.Insert(dummyKey, dummyVal)

	actualForwardMap := actual.GetForwardMap()
	eq := reflect.DeepEqual(actualForwardMap, forwardMap)
	assert.True(t, eq, "Forward maps should be equal")
}

func TestBiMap_GetInverseMap(t *testing.T) {
	actual := NewBiMap[string, int]()
	dummyKey := "Dummy key"
	dummyVal := 42

	inverseMap := make(map[int]string)
	inverseMap[dummyVal] = dummyKey

	actual.Insert(dummyKey, dummyVal)

	actualInverseMap := actual.GetInverseMap()
	eq := reflect.DeepEqual(actualInverseMap, inverseMap)
	assert.True(t, eq, "Inverse maps should be equal")
}
