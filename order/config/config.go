package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	AppName     string `envconfig:"APP_NAME" default:"order-service"`
	Host        string `envconfig:"HOST" default:"order-service"`
	Port        string `envconfig:"PORT" default:"3000"`
	Environment string `envconfig:"ENVIRONMENT" default:"dev"`
	AuthService AuthService
	KafkaConfig Kafka
}

type AuthService struct {
	BaseURL string `envconfig:"AUTH_SERVICE_BASE_URL" default:"http://0.0.0.0:5001"`
}

type Kafka struct {
	Servers string `envconfig:"KAFKA_SERVERS" default:"broker:9093,broker:9093,broker:9093"`
	Timeout int    `envconfig:"KAFKA_CONN_TIMEOUT" default:"6000"`
}

func Load() (*Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}

	return &config, nil
}
