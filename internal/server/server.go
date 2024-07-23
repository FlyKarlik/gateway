package server

import (
	"gateway/config"
	"gateway/internal/client"
	"github.com/hashicorp/go-hclog"
	"go.opentelemetry.io/otel/trace"
	"sync"
)

func RunServer(wg *sync.WaitGroup, trace trace.Tracer, cfg *config.Config, responseHash *client.MessageHash) {
	log := hclog.Default()
	if cfg == nil {
		log.Error("RunServer error", "error", "config is nil")
		return
	}
	defer wg.Done()

	producer := client.NewServiceProducer(log, cfg)

	r := NewRouter(trace, cfg, producer, responseHash)
	err := r.Run(cfg.ServerHost)
	log.Error("can not serve", "error", err)
}

func RunKafkaConsumer(wg *sync.WaitGroup, cfg *config.Config, responseHash *client.MessageHash) {
	log := hclog.Default()
	mapsConsumer := client.NewMapsResponseConsumer(log, cfg, responseHash)

	if cfg == nil {
		log.Error("RunServer error", "error", "config is nil")
		return
	}
	defer wg.Done()

	mapsConsumer.Run()
}
