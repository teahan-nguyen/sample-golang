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
	Role      string `mapstructure:"ROLE"`
	DbName    string `mapstructure:"DBNAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		return
	}
	return
}
