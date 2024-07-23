package controllers

import (
	"comet/utils"
	"context"
	"gateway/internal/client"
	"gateway/internal/controllers/models"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	pb "protos/maps"
	"time"
)

func (cn *Controllers) AddLayer(c *gin.Context) {
	log := hclog.Default()

	var data models.Layer

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "AddLayer", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controllers.AddLayer] utils.HandleRequest", "error", err)
		return
	}

	dataBytes, err := proto.Marshal(&pb.MLayer{
		LayerType:    data.LayerType,
		TableId:      data.TableID,
		Name:         data.Name,
		CreateUserId: "Odil",
	})
	if err != nil {
		log.Error("[controllers.AddLayer] json.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   dataBytes,
		Key:     []byte{utils.AddLayerRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.AddLayerRequestPartition},
	})

	if err != nil {
		log.Error("[controllers.AddLayer] cn.Producer.SendMessage", "error", err)
		return
	}

	var model pb.LayerMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) Layer(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "Layer", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.Layer] utils.HandleRequest", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	layerID := c.Query("id")
	if len(layerID) < 1 {
		log.Error("[controller.Layer] c.Query(id)", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MLayer{Id: layerID})
	if err != nil {
		log.Error("[controller.Layer] proto.Marshal", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.LayerRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.LayerRequestPartition},
	})

	var model pb.LayerMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) DeleteLayer(c *gin.Context) {
	var data models.Layer
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "DeleteLayer", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.DeleteLayer] utils.HandleRequest", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MLayer{Id: data.ID})
	if err != nil {
		log.Error("[controller.DeleteLayer] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.DeleteLayerRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.DeleteLayerRequestPartition},
	})

	var model pb.LayerMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) Layers(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "Layer", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.Layers] utils.HandleRequest", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   []byte{},
		Key:     []byte{utils.LayersRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.LayersRequestPartition},
	})

	var model pb.LayersMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) EditLayer(c *gin.Context) {
	var data models.Layer
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "EditLayer", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.EditLayer] utils.HandleRequest", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MLayer{
		Id:        data.ID,
		Name:      data.Name,
		TableId:   data.TableID,
		LayerType: data.LayerType,
		UpdatedAt: time.Now().String(),
	})
	if err != nil {
		log.Error("[controller.EditLayer] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.EditLayerRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.EditLayerRequestPartition},
	})

	var model pb.LayerMessage
	cn.waitResponse(cc, c, id, &model)

	return
}
