package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStrFromCfg(t *testing.T) {
	t.Run("testGetStrFromCfgValid", testGetStrFromCfgValid)
	t.Run("testGetStrFromCfgMissing", testGetStrFromCfgMissing)
	t.Run("testGetStrFromCfgNotStr", testGetStrFromCfgNotStr)
	t.Run("testGetStrFromCfgEmptyStr", testGetStrFromCfgEmptyStr)
}

func testGetStrFromCfgValid(t *testing.T) {
	cfg := ComponentConfig{
		"foo": "bar",
	}

	res, err := cfg.GetStr("foo")

	assert.NoError(t, err)
	assert.Equal(t, "bar", res)
}

func testGetStrFromCfgMissing(t *testing.T) {
	cfg := ComponentConfig{}

	res, err := cfg.GetStr("foo")

	assert.Error(t, err)
	assert.Empty(t, res)
}

func testGetStrFromCfgNotStr(t *testing.T) {
	cfg := ComponentConfig{
		"foo": map[string]string{"bar": "bar"},
	}

	res, err := cfg.GetStr("foo")

	assert.Error(t, err)
	assert.Empty(t, res)
}

func testGetStrFromCfgEmptyStr(t *testing.T) {
	cfg := ComponentConfig{
		"foo": "",
	}

	res, err := cfg.GetStr("foo")

	assert.NoError(t, err)
	assert.Empty(t, res)
}

func TestExtend(t *testing.T) {
	t.Run("testExtendWithEmpty", testExtendWithEmpty)
	t.Run("testExtendOnEmpty", testExtendOnEmpty)
	t.Run("testExtendOverwrite", testExtendOverwrite)
	t.Run("testExtendOverwriteComplex", testExtendOverwriteComplex)
}

func testExtendWithEmpty(t *testing.T) {
	base := ComponentConfig{
		"foo": "bar",
	}

	extension := ComponentConfig{}

	assert.Equal(t, base, base.Extend(extension))
}

func testExtendOnEmpty(t *testing.T) {
	base := ComponentConfig{}

	extension := ComponentConfig{
		"foo": "bar",
	}

	assert.Equal(t, extension, base.Extend(extension))
}

func testExtendOverwrite(t *testing.T) {
	base := ComponentConfig{
		"foo": "bar1",
	}

	extension := ComponentConfig{
		"foo": "bar",
	}

	assert.Equal(t, "bar", base.Extend(extension)["foo"].(string))
}

func testExtendOverwriteComplex(t *testing.T) {
	base := ComponentConfig{
		"foo": map[string]string{
			"foo1": "bar1",
		},
	}

	extension := ComponentConfig{
		"foo": "bar",
	}

	assert.Equal(t, "bar", base.Extend(extension)["foo"].(string))
}