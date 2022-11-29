package _yaml

import "github.com/spf13/viper"

// LoadCoreConfig
func LoadCoreConfig(config interface{}) {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(config)
	if err != nil {
		panic(err)
	}
}
