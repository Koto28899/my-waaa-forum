package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Dsn                           string        `mapsturcture:"dsn"`
	ServerAddress                 string        `mapsturcture:"serverAddress"`
	TokenSymmetricKey             string        `mapsturcture:"tokenSymmetricKey"`
	AccessTokenDuration           time.Duration `mapsturcture:"accessTokenDuration"`
	SessionDuration               time.Duration `mapstructure:"sessionDuration"`
	AccessControlAllowOrigin      string        `mapstructure:"accessControlAllowOrigin"`
	AccessControlAllowMethods     string        `mapstructure:"accessControlAllowMethods"`
	AccessControlAllowCredentials string        `mapstructure:"accessControlAllowCredentials"`
	CookieName                    string        `mapstructure:"cookieName"`
	CookiePath                    string        `mapstructure:"cookiePath"`
	CookieDomain                  string        `mapstructure:"cookieDomain"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("serverConfig")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
