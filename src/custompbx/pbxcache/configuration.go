package pbxcache

import (
	"custompbx/altStruct"
)

var ConfigurationCache altStruct.Configurations

func GetConfigs(s []interface{}) *altStruct.Configurations {
	conf := &altStruct.Configurations{}
	for _, c := range s {
		config, ok := c.(altStruct.ConfigurationsList)
		if !ok {
			continue
		}
		subCached := ConfigurationCache.GetConfiguration(config.Name)
		if subCached != nil {
			config.Loaded = subCached.Loaded
		}
		conf.FillConfigurations(&config)
	}
	return conf
}
