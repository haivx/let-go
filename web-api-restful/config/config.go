package config

import "github.com/spf13/viper"

type Config struct {
	DB_SOURCE            string `mapstructure:"DB_SOURCE"`
	MIGRATION_URL        string `mapstructure:"MIGRATION_URL"`
	POSTGRES_USER        string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PORT        string `mapstructure:"POSTGRES_PORT"`
	POSTGRES_DB          string `mapstructure:"POSTGRES_DB"`
	POSTGRES_PASSWORD    string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_HOST        string `mapstructure:"POSTGRES_HOST"`
	JWT_SECRET_KEY       string `mapstructure:"JWT_SECRET_KEY"`
	MAIL_HOST            string `mapstructure:"MAIL_HOST"`
	MAIL_SENDER_USERNAME string `mapstructure:"MAIL_SENDER_USERNAME"`
	MAIL_SENDER_PASSWORD string `mapstructure:"MAIL_SENDER_PASSWORD"`
	MAIL_PORT            int    `mapstructure:"MAIL_PORT"`
	MAIL_SUBJECT         string `mapstructure:"MAIL_SUBJECT"`
	MAIL_TO              string `mapstructure:"MAIL_TO"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
