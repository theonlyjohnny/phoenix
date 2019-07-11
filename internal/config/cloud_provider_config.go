package config

var validClouds = []string{"ec2"}

//CloudProviderConfig is an arbitrary JSON interface for use by individual clouds
type CloudProviderConfig map[string]interface{}

func (cpc CloudProviderConfig) Extend(override CloudProviderConfig) CloudProviderConfig {
	for k, v := range override {
		cpc[k] = v
	}
	return cpc
}
