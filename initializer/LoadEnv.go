package initializer

import "github.com/spf13/viper"

type Config struct {
	ClientId     string `mapstructure:"CLIENT_ID"`
	ClientSecret string `mapstructure:"CLIENT_SECRET"`
	Issuer       string `mapstructure:"ISSUER"`
	SpaClientId  string `mapstructure:"SPA_CLIENT_ID"`

	ClientOrigin  string `mapstructure:"CLIENT_ORIGIN"`
	ClientOrigin2 string `mapstructure:"CLIENT_ORIGIN2"`
	ServerPort    string `mapstructure:"PORT"`
	UriAddress    string `mapstructure:"URI_ADDRESS"`

	SecretKey string `mapstructure:"SECRET_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
