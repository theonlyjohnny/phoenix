package config_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theonlyjohnny/phoenix/internal/config"
)

var (
	fileBase = fmt.Sprintf("%s/config_test", os.Getenv("TEST_OUTPUT"))
)

func TestConsistentDefault(t *testing.T) {
	assert.Equal(t, config.DefaultConfig(), config.DefaultConfig())
}

func TestReadConfigFromFs(t *testing.T) {
	if err := os.MkdirAll(fileBase, 0755); err != nil {
		t.Fatal(fmt.Sprintf("Could not create %s: %s", fileBase, err.Error()))
	}
	t.Run("testReadInvalidFile", testReadInvalidFile)
	t.Run("testReadInvalidJSON", testReadInvalidJSON)
	t.Run("testReadInvalidCloud", testReadInvalidCloud)
	t.Run("testReadValidCloud", testReadValidCloud)
	t.Run("testReadInvalidStorage", testReadInvalidStorage)
	t.Run("testReadValidStorage", testReadValidStorage)
}

func testReadInvalidFile(t *testing.T) {
	assert.Equal(t, config.DefaultConfig(), config.ReadConfigFromFs("/"))
}

func testReadInvalidJSON(t *testing.T) {
	path := fmt.Sprintf("%s/invalidjson.json", fileBase)
	if err := ioutil.WriteFile(path, []byte("{,}"), 0755); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, config.DefaultConfig(), config.ReadConfigFromFs(path))

}

func testReadInvalidCloud(t *testing.T) {
	path := fmt.Sprintf("%s/invalidcloud.json", fileBase)
	contents := "{\"cloud_type\": \"invalid\"}"
	if err := ioutil.WriteFile(path, []byte(contents), 0755); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, config.DefaultConfig(), config.ReadConfigFromFs(path))

}

func testReadValidCloud(t *testing.T) {
	expectedCloudType := "ec2"

	path := fmt.Sprintf("%s/validcloud.json", fileBase)
	contents := fmt.Sprintf("{\"cloud_type\": \"%s\"}", expectedCloudType)
	if err := ioutil.WriteFile(path, []byte(contents), 0755); err != nil {
		t.Fatal(err)
	}

	expected := config.DefaultConfig()
	expected.CloudType = expectedCloudType

	assert.Equal(t, expected, config.ReadConfigFromFs(path))

}

func testReadInvalidStorage(t *testing.T) {
	path := fmt.Sprintf("%s/invalidstorage.json", fileBase)
	contents := "{\"storage_type\": \"invalid\"}"
	if err := ioutil.WriteFile(path, []byte(contents), 0755); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, config.DefaultConfig(), config.ReadConfigFromFs(path))
}

func testReadValidStorage(t *testing.T) {
	expectedStorageType := "redis"

	path := fmt.Sprintf("%s/validstorage.json", fileBase)
	contents := fmt.Sprintf("{\"storage_type\": \"%s\"}", expectedStorageType)
	if err := ioutil.WriteFile(path, []byte(contents), 0755); err != nil {
		t.Fatal(err)
	}

	expected := config.DefaultConfig()
	expected.StorageType = expectedStorageType

	assert.Equal(t, expected, config.ReadConfigFromFs(path))

}
