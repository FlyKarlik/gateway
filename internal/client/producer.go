package client

import (
	"gateway/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/hashicorp/go-hclog"
)

type ServiceProducer struct {
	log      hclog.Logger
	cfg      *config.Config
	Producer *kafka.Producer
}

// NewServiceProducer constructor
func NewServiceProducer(log hclog.Logger, cfg *config.Config) *ServiceProducer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.KafkaBrokers,
		"client.id":         cfg.ServiceName,
	})

	if err != nil {
		log.Error("[client.GetNewKafkaProducer] kafka.NewProducer failed", "error", err)
	}

	return &ServiceProducer{log: log, cfg: cfg, Producer: producer}
}

func (p *ServiceProducer) SendMessage(msg *kafka.Message) error {
	return p.Producer.Produce(msg, nil)
}

func (p *ServiceProducer) Close() {
	p.Producer.Close()
}
