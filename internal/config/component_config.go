package config

import "fmt"

//ComponentConfig is an arbitrary JSON interface for use by individual clouds
type ComponentConfig map[string]interface{}

func (cpc ComponentConfig) Extend(override ComponentConfig) ComponentConfig {
	for k, v := range override {
		cpc[k] = v
	}
	return cpc
}

func (cpc ComponentConfig) GetStr(k string) (string, error) {
	var res string

	if interf, ok := cpc[k]; ok {
		if cpcStr, ok := interf.(string); ok {
			res = cpcStr
		} else {
			return res, fmt.Errorf("config.%s is not a string", k)
		}
	} else {
		return "", fmt.Errorf("config.%s not found", k)
	}
	return res, nil
}

func (cpc ComponentConfig) GetInt(k string) (int, error) {
	var res int

	if interf, ok := cpc[k]; ok {
		if cpcStr, ok := interf.(int); ok {
			res = cpcStr
		} else {
			return res, fmt.Errorf("config.%s is not a int", k)
		}
	} else {
		return res, fmt.Errorf("config.%s is a required parameter", k)
	}
	return res, nil
}

func (cpc ComponentConfig) GetNestedConfigComponent(k string) (ComponentConfig, error) {
	var res ComponentConfig

	if interf, ok := cpc[k]; ok {
		if cpcStr, ok := interf.(ComponentConfig); ok {
			res = cpcStr
		} else {
			return res, fmt.Errorf("config.%s is not a nested config", k)
		}
	} else {
		return nil, fmt.Errorf("config.%s not found", k)
	}
	return res, nil
}
