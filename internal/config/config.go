package config

import "github.com/spf13/viper"

type Config struct {
	Port             string `mapstructure:"PORT"`
	DBUrl            string `mapstructure:"DB_URL"`
	PassengerSvcUrl  string `mapstructure:"PASSENGER_SVC_URL"`
	DriverSvcUrl     string `mapstructure:"DRIVER_SVC_URL"`
	CallCenterSvcUrl string `mapstructure:"CALLCENTER_SVC_URL"`
	JWTSecretKey     string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./internal/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
