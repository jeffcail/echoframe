package _yaml

import "github.com/spf13/viper"

// LoadCoreConfig
func LoadCoreConfig(path string, config interface{}) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(config)
	if err != nil {
		panic(err)
	}
}
