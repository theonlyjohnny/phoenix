package config

import "fmt"

var validClouds = []string{"ec2"}

//CloudProviderConfig is an arbitrary JSON interface for use by individual clouds
type CloudProviderConfig map[string]interface{}

func (cpc CloudProviderConfig) Extend(override CloudProviderConfig) CloudProviderConfig {
	for k, v := range override {
		cpc[k] = v
	}
	return cpc
}

func (cpc CloudProviderConfig) GetStr(k string) (string, error) {
	var res string

	if interf, ok := cpc[k]; ok {
		if cpcStr, ok := interf.(string); ok {
			if cpcStr == "" {
				return res, fmt.Errorf("cloud_config.%s cannot be an empty string", k)
			}
			res = cpcStr
		} else {
			return res, fmt.Errorf("cloud_config.%s is not a string", k)
		}
	} else {
		return res, fmt.Errorf("cloud_config.%s is a required parameter", k)
	}
	return res, nil
}
