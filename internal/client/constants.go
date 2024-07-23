package client

import "time"

const (
	minBytes               = 10e3 // 10KB
	maxBytes               = 10e6 // 10MB
	queueCapacity          = 100
	heartbeatInterval      = 3 * time.Second
	commitInterval         = 0
	partitionWatchInterval = 5 * time.Second
	MaxAttempts            = 3
	DialTimeout            = 10 * time.Second

	writerReadTimeout  = 10 * time.Second
	writerWriteTimeout = 10 * time.Second
	writerRequiredAcks = -1
	writerMaxAttempts  = 3

	createLayerTopic   = "create-layer"
	createLayerWorkers = 3
	updateLayerTopic   = "update-layer"
	updateLayerWorkers = 3

	deadLetterQueueTopic = "dead-letter-queue"

	productsGroupID = "maps_group"

	AddLayerRequest    = "add_layer_request"
	AddLayerResponse   = "add_layer_response"
	LayerRequest       = "layer_request"
	LayerResponse      = "layer_response"
	EditLayerRequest   = "edit_layer_request"
	EditLayerResponse  = "edit_layer_response"
	DeleteLayerRequest = "delete_layer_request"
)
