package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStrFromCfg(t *testing.T) {
	t.Run("testGetStrFromCfgValid", testGetStrFromCfgValid)
	t.Run("testGetStrFromCfgNotFound", testGetStrFromCfgNotFound)
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

func testGetStrFromCfgNotFound(t *testing.T) {
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

func TestGetIntFromCfg(t *testing.T) {
	t.Run("testGetIntFromCfgValid", testGetIntFromCfgValid)
	t.Run("testGetIntFromCfgNotFound", testGetIntFromCfgNotFound)
	t.Run("testGetIntFromCfgNotInt", testGetIntFromCfgNotInt)
	t.Run("testGetIntFromCfgEmptyInt", testGetIntFromCfgEmptyInt)
}

func testGetIntFromCfgValid(t *testing.T) {
	cfg := ComponentConfig{
		"foo": 12,
	}

	res, err := cfg.GetInt("foo")

	assert.NoError(t, err)
	assert.Equal(t, 12, res)
}

func testGetIntFromCfgNotFound(t *testing.T) {
	cfg := ComponentConfig{}

	res, err := cfg.GetInt("foo")

	assert.Error(t, err)
	assert.Empty(t, res)
}

func testGetIntFromCfgNotInt(t *testing.T) {
	cfg := ComponentConfig{
		"foo": map[string]string{"bar": "bar"},
	}

	res, err := cfg.GetInt("foo")

	assert.Error(t, err)
	assert.Empty(t, res)
}

func testGetIntFromCfgEmptyInt(t *testing.T) {
	cfg := ComponentConfig{
		"foo": 0,
	}

	res, err := cfg.GetInt("foo")

	assert.NoError(t, err)
	assert.Empty(t, res)
}

func TestGetConfigComponentFromCfg(t *testing.T) {
	t.Run("testGetConfigComponentFromCfgValid", testGetConfigComponentFromCfgValid)
	t.Run("testGetConfigComponentFromCfgNotFound", testGetConfigComponentFromCfgNotFound)
	t.Run("testGetConfigComponentFromCfgNotConfigComponent", testGetConfigComponentFromCfgNotConfigComponent)
	t.Run("testGetConfigComponentFromCfgEmptyConfigComponent", testGetConfigComponentFromCfgEmptyConfigComponent)
}

func testGetConfigComponentFromCfgValid(t *testing.T) {
	nested := ComponentConfig{
		"foo":  "bar",
		"foo2": 12,
	}
	cfg := ComponentConfig{
		"nested": nested,
	}

	res, err := cfg.GetNestedConfigComponent("nested")

	assert.NoError(t, err)
	assert.Equal(t, nested, res)
}

func testGetConfigComponentFromCfgNotFound(t *testing.T) {
	cfg := ComponentConfig{}

	res, err := cfg.GetNestedConfigComponent("foo")

	assert.Error(t, err)
	assert.Empty(t, res)
}

func testGetConfigComponentFromCfgNotConfigComponent(t *testing.T) {
	cfg := ComponentConfig{
		"foo": []string{"bar", "bar"},
	}

	res, err := cfg.GetNestedConfigComponent("foo")

	assert.Error(t, err)
	assert.Empty(t, res)
}

func testGetConfigComponentFromCfgEmptyConfigComponent(t *testing.T) {
	cfg := ComponentConfig{
		"foo": ComponentConfig{},
	}

	res, err := cfg.GetNestedConfigComponent("foo")

	assert.NoError(t, err)
	assert.Empty(t, res)
}
