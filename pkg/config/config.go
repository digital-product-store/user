package config

import (
	"time"

	"github.com/caarlos0/env/v9"
)

type Config struct {
	ListenAddr    string        `env:"LISTEN_ADDR" envDefault:":8000"`
	JWTPrivateKey string        `env:"JWT_PRIVATE_KEY" envDefault:"keys/private.pem"`
	JWTIssuer     string        `env:"JWT_ISSUER" envDefault:"user@a-digitalstore.test.local"`
	JWTSubject    string        `env:"JWT_SUBJECT" envDefault:"user@a-digitalstore.test.local"`
	JWTTTL        time.Duration `env:"JWT_TTL" envDefault:"24h"`
}

func LoadConfig() (*Config, error) {
	conf := &Config{}
	err := env.Parse(conf)
	return conf, err
}
