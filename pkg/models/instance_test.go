package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theonlyjohnny/phoenix/pkg/models"
)

const (
	testInstanceName = "test_instance_name"
)

func TestLocationString(t *testing.T) {
	assert.NotEmpty(t, models.Location{}.String())
}

func TestInstanceString(t *testing.T) {
	assert.NotEmpty(t, models.Instance{}.String())
}

func TestNewInstance(t *testing.T) {
	t.Run("testNewInstance", testNewInstance)
	t.Run("testNewInstanceIDUnique", testNewInstanceIDUnique)
}

func testNewInstance(t *testing.T) {
	instance := models.NewInstance(testInstanceName)
	assert.Equal(t, testInstanceName, instance.Name)
	assert.NotEmpty(t, instance.PhoenixID)
}

func testNewInstanceIDUnique(t *testing.T) {
	first := models.NewInstance(testInstanceName)
	assert.Equal(t, testInstanceName, first.Name)
	assert.NotEmpty(t, first.PhoenixID)

	second := models.NewInstance(testInstanceName)
	assert.Equal(t, testInstanceName, first.Name)
	assert.NotEmpty(t, second.PhoenixID)

	assert.NotEqual(t, first.PhoenixID, second.PhoenixID)
}
