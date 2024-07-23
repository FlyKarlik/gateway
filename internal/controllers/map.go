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
	"net/http"
	pb "protos/maps"
	"time"
)

func (cn *Controllers) AddMap(c *gin.Context) {
	log := hclog.Default()

	var data models.Map

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "AddMap", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controllers.AddMap] utils.HandleRequest", "error", err)
		return
	}

	dataBytes, err := proto.Marshal(&pb.MMap{
		Name:         data.Name,
		Picture:      data.Picture,
		Describe:     data.Describe,
		Active:       data.Active,
		CreateUserId: "Odil",
		CreateUserIp: c.ClientIP(),
	})
	if err != nil {
		log.Error("[controllers.AddMap] json.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   dataBytes,
		Key:     []byte{utils.AddMapRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.AddMapRequestPartition},
	})

	if err != nil {
		log.Error("[controllers.AddMap] cn.Producer.SendMessage", "error", err)
		return
	}

	var model pb.MapMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) Map(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "Map", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.Map] utils.HandleRequest", "error", err)
		return
	}

	mapID := c.Query("id")
	if len(mapID) < 1 {
		c.Set("message", "not correct ID")
		c.Set("code", http.StatusBadRequest)
		c.Set("status", "failed")

		log.Error("[controller.Map] proto.Marshal", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MMap{Id: mapID})
	if err != nil {
		log.Error("[controller.Map] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.MapRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.MapRequestPartition},
	})

	var model pb.MapMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) DeleteMap(c *gin.Context) {
	var data models.Map
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "DeleteMap", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.DeleteMap] utils.HandleRequest", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MMap{Id: data.ID})
	if err != nil {
		log.Error("[controller.DeleteMap] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.DeleteMapRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.DeleteMapRequestPartition},
	})

	var model pb.MapMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) Maps(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "Maps", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.Maps] utils.HandleRequest", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   []byte{},
		Key:     []byte{utils.MapsRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.MapsRequestPartition},
	})

	var model pb.MapsMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) EditMap(c *gin.Context) {
	var data models.Map
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "EditMap", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.EditMap] utils.HandleRequest", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MMap{
		Id:           data.ID,
		Name:         data.Name,
		Picture:      data.Picture,
		Describe:     data.Describe,
		Active:       data.Active,
		UpdateUserId: "Odil",
		UpdateUserIp: c.ClientIP(),
	})
	if err != nil {
		log.Error("[controller.EditMap] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.EditMapRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.EditMapRequestPartition},
	})

	var model pb.MapMessage
	cn.waitResponse(cc, c, id, &model)

	return
}
