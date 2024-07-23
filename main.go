package main

import (
	"comet/utils"
	"fmt"
	"gateway/config"
	"gateway/internal/client"
	"gateway/internal/prometheus"
	"gateway/internal/server"
	"github.com/getsentry/sentry-go"
	"sync"
	"time"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	cfg, err := config.InitConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	err = utils.InitSentry(cfg.SentryDSN)
	if err != nil {
		return fmt.Errorf("failed init sentry: %w", err)
	}
	defer sentry.Flush(2 * time.Second)

	trace, err := utils.InitTracerProvider(cfg.JaegerHost, "gateway")
	if err != nil {
		return fmt.Errorf("filed init tracer: %w", err)
	}

	responseHash := client.NewMessageHash()

	var wg sync.WaitGroup
	wg.Add(3)

	go server.RunServer(&wg, trace, cfg, responseHash)
	go server.RunKafkaConsumer(&wg, cfg, responseHash)
	go prometheus.RunPrometheus(&wg, cfg.PrometheusHost)

	wg.Wait()

	return nil
}
