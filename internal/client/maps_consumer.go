package client

import (
	"gateway/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/hashicorp/go-hclog"
	"time"
)

type MapsResponseConsumer struct {
	Logger   hclog.Logger
	cfg      *config.Config
	Msgs     *MessageHash
	Consumer *kafka.Consumer
}

func NewMapsResponseConsumer(log hclog.Logger, cfg *config.Config, responseHash *MessageHash) *MapsResponseConsumer {
	k, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.KafkaBrokers,
		"group.id":          cfg.ServiceName,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Error("[client.GetLayerCreateResponse] kafka.NewConsumer failed", "error", err)
	}

	err = k.SubscribeTopics([]string{cfg.KafkaMapsResponseTopic}, nil)
	if err != nil {
		log.Error("[client.GetLayerCreateResponse] k.SubscribeTopics failed", "error", err)
	}

	return &MapsResponseConsumer{
		Logger:   hclog.Default(),
		cfg:      cfg,
		Msgs:     responseHash,
		Consumer: k,
	}
}

func (c *MapsResponseConsumer) Run() {
	go func() {
		for {
			ev, err := c.Consumer.ReadMessage(5 * time.Millisecond)
			if err != nil {
				continue
			}

			c.Logger.Info("[Consumer worker]", "topic", *ev.TopicPartition.Topic, "partition", ev.TopicPartition.Partition, "key", ev.Key, "hash_store_cap", len(c.Msgs.Hash))

			id := ev.Headers[0].Key
			c.Msgs.Add(id, Node{Message: ev.Value, P: ev.TopicPartition.Partition})
		}
	}()
}
