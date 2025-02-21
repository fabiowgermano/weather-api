package configs

import "github.com/spf13/viper"

type conf struct {
	LocationClientUrl string `mapstructure:"LOCATION_URL"`
	WeatherClientUrl  string `mapstructure:"WEATHER_URL"`
	WeatherClientKey  string `mapstructure:"WEATHER_KEY"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.SetConfigFile(path + "/.env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
