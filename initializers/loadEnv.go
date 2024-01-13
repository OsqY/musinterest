package initializers

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"DB_HOST"`
	DBUsername     string `mapstructure:"DB_USER"`
	DBUserPassword string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	DBPort         string `mapstructure:"DB_PORT"`

	Auth0Domain       string `mapstructure:"AUTH0_DOMAIN"`
	Auth0ClientID     string `mapstructure:"AUTH0_CLIENT_ID"`
	Auth0ClientSecret string `mapstructure:"AUTH0_CLIENT_SECRET"`
	Auth0CallbackURL  string `mapstructure:"AUTH0_CALLBACK_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env.local")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
