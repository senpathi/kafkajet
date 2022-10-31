package repository

import (
	"github.com/caarlos0/env/v6"
	"log"
)

var DBConf DBConfig

type DBConfig struct {
	Database      string `env:"DB_NAME" envDefault:"kafka_jet"`
	User          string `env:"DB_USER" envDefault:"root"`
	Password      string `env:"DB_PASSWORD" envDefault:"root"`
	Address       string `env:"DB_URL" envDefault:"localhost:27017"`
	MaxConnection int    `env:"MAX_DB_CONNECTION" envDefault:"10"`
}

func (d *DBConfig) Load() {
	err := env.Parse(&DBConf)
	if err != nil {
		log.Fatalln("error parsing app configuration", err)
	}

	*d = DBConf
}
