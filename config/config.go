package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var cfg *Config

type Config struct {
	ServiceName                    string
	ServerHost                     string
	SentryDSN                      string
	JaegerHost                     string
	PrometheusHost                 string
	PrivateKeyLocation             string
	PublicKeyLocation              string
	AccountServiceHost             string
	MapsServiceHost                string
	SurveyServiceHost              string
	CustomLayersServiceHost        string
	KafkaBrokers                   string
	KafkaAutoOffsetRest            string
	KafkaMapsResponseTopic         string
	KafkaMapsRequestTopic          string
	KafkaAccountResponseTopic      string
	KafkaAccountRequestTopic       string
	KafkaSurveyResponseTopic       string
	KafkaSurveyRequestTopic        string
	KafkaCustomLayersResponseTopic string
	KafkaCustomLayersRequestTopic  string
}

func InitConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("initConfig godotenv load failed")
	}

	cfg = &Config{
		ServiceName:                    os.Getenv("SERVICE_NAME"),
		ServerHost:                     os.Getenv("SERVER_HOST"),
		AccountServiceHost:             os.Getenv("ACCOUNT_SERVICE_HOST"),
		MapsServiceHost:                os.Getenv("MAPS_SERVICE_HOST"),
		SurveyServiceHost:              os.Getenv("SURVEY_SERVICE_HOST"),
		CustomLayersServiceHost:        os.Getenv("CUSTOM_LAYERS_SERVICE_HOST"),
		SentryDSN:                      os.Getenv("SENTRY_DNS"),
		JaegerHost:                     os.Getenv("JAEGER_HOST"),
		PrometheusHost:                 os.Getenv("PROMETHEUS_HOST"),
		PrivateKeyLocation:             os.Getenv("PRIVATE_KEY_LOCATION"),
		PublicKeyLocation:              os.Getenv("PUBLIC_KEY_LOCATION"),
		KafkaBrokers:                   os.Getenv("KAFKA_BROKERS"),
		KafkaAutoOffsetRest:            os.Getenv("KAFKA_AUTO_OFFSET_RESET"),
		KafkaMapsRequestTopic:          os.Getenv("KAFKA_MAPS_REQUEST_TOPIC"),
		KafkaMapsResponseTopic:         os.Getenv("KAFKA_MAPS_RESPONSE_TOPIC"),
		KafkaAccountRequestTopic:       os.Getenv("KAFKA_ACCOUNT_REQUEST_TOPIC"),
		KafkaAccountResponseTopic:      os.Getenv("KAFKA_ACCOUNT_RESPONSE_TOPIC"),
		KafkaSurveyRequestTopic:        os.Getenv("KAFKA_SURVEY_REQUEST_TOPIC"),
		KafkaSurveyResponseTopic:       os.Getenv("KAFKA_SURVEY_RESPONSE_TOPIC"),
		KafkaCustomLayersRequestTopic:  os.Getenv("KAFKA_CUSTOM_LAYERS_REQUEST_TOPIC"),
		KafkaCustomLayersResponseTopic: os.Getenv("KAFKA_CUSTOM_LAYERS_RESPONSE_TOPIC"),
	}

	return cfg, nil
}
