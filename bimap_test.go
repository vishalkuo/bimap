package bimap

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewBiMap(t *testing.T) {
	biMap := NewBiMap()
	expected := BiMap{forward:make(map[string]string), inverse:make(map[string]string)}
	assert.Equal(t, *biMap, expected, "They should be equal")
}
