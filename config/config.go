package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Pg struct {
			Host                  string        `mapstructure:"HOST"`
			Port                  int           `mapstructure:"PORT"`
			Dbname                string        `mapstructure:"DBNAME"`
			User                  string        `mapstructure:"USER"`
			Password              string        `mapstructure:"PASSWORD"`
			SslMode               string        `mapstructure:"SSLMODE"`
			MaxOpenConnection     int           `mapstructure:"MAX_OPEN_CONNECTION"`
			MaxIdleConnection     int           `mapstructure:"MAX_IDLE_CONNECTION"`
			MaxConnectionLifetime time.Duration `mapstructure:"MAX_CONNECTION_LIFETIME"`
		} `mapstructure:"PG"`
	}
}

func InitConfig() (config *Config) {
	log.Println("reading config")

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("error reading .env file", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln("failed unmarshaling config", err)
	}

	log.Println("config ready")

	return
}
