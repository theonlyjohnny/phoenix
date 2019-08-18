package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testInstanceName = "test_instance_name"
)

func TestLocationString(t *testing.T) {
	assert.NotEmpty(t, Location{}.String())
}

func TestInstanceString(t *testing.T) {
	assert.NotEmpty(t, Instance{}.String())
}

func TestNewInstance(t *testing.T) {
	t.Run("testNewInstance", testNewInstance)
	t.Run("testNewInstanceIDUnique", testNewInstanceIDUnique)
}

func testNewInstance(t *testing.T) {
	instance := NewInstance(testInstanceName)
	assert.Equal(t, testInstanceName, instance.Name)
	assert.NotEmpty(t, instance.PhoenixID)
}

func testNewInstanceIDUnique(t *testing.T) {
	first := NewInstance(testInstanceName)
	assert.Equal(t, testInstanceName, first.Name)
	assert.NotEmpty(t, first.PhoenixID)

	second := NewInstance(testInstanceName)
	assert.Equal(t, testInstanceName, first.Name)
	assert.NotEmpty(t, second.PhoenixID)

	assert.NotEqual(t, first.PhoenixID, second.PhoenixID)
}
