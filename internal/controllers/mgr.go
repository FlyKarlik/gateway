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

func (cn *Controllers) AddMapGroupRelation(c *gin.Context) {
	log := hclog.Default()

	var data models.MapGroupRelation

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "AddMapGroupRelation", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controllers.AddMapGroupRelation] utils.HandleRequest", "error", err)
		return
	}

	dataBytes, err := proto.Marshal(&pb.MGRelation{
		GroupId:    data.GroupID,
		MapId:      data.MapID,
		GroupOrder: data.GroupOrder,
	})
	if err != nil {
		log.Error("[controllers.AddMapGroupRelation] json.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   dataBytes,
		Key:     []byte{utils.AddMapGroupRelationRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.AddMapGroupRelationRequestPartition},
	})

	if err != nil {
		log.Error("[controllers.AddMapGroupRelation] cn.Producer.SendMessage", "error", err)
		return
	}

	var model pb.MGRelationMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) MapGroupRelations(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "MapGroupRelations", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.MapGroupRelations] utils.HandleRequest", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   []byte{},
		Key:     []byte{utils.MapGroupRelationsRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.MapGroupRelationsRequestPartition},
	})

	var model pb.MGRelationsMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) DeleteMapGroupRelation(c *gin.Context) {
	var data models.MapGroupRelation
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "DeleteMapGroupRelation", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.DeleteMapGroupRelation] utils.HandleRequest", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MGRelation{Id: data.ID})
	if err != nil {
		log.Error("[controller.DeleteMapGroupRelation] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.DeleteMapGroupRelationRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.DeleteMapGroupRelationRequestPartition},
	})

	var model pb.MGRelationMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) MapRelationGroups(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "MapRelationGroups", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.MapRelationGroups] utils.HandleRequest", "error", err)
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
		log.Error("[controller.MapRelationGroups] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.MapRelationGroupsRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.MapRelationGroupsRequestPartition},
	})

	var model pb.GroupsMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) GroupRelationMaps(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "GroupRelationMaps", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.GroupRelationMaps] utils.HandleRequest", "error", err)
		return
	}

	groupID := c.Query("id")
	if len(groupID) < 1 {
		c.Set("message", "not correct ID")
		c.Set("code", http.StatusBadRequest)
		c.Set("status", "failed")

		log.Error("[controller.Map] proto.Marshal", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MGroup{Id: groupID})
	if err != nil {
		log.Error("[controller.GroupRelationMaps] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.GroupRelationMapsRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.GroupRelationMapsRequestPartition},
	})

	var model pb.MapsMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) MapGroupOrderUp(c *gin.Context) {
	var data models.MapGroupRelation
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "MapGroupOrderUp", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.MapGroupOrderUp] utils.HandleRequest", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MGRelation{Id: data.ID})
	if err != nil {
		log.Error("[controller.MapGroupOrderUp] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.MapGroupOrderUpRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.MapGroupOrderUpRequestPartition},
	})

	var model pb.MGRelationMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) MapGroupOrderDown(c *gin.Context) {
	var data models.MapGroupRelation
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "MapGroupOrderDown", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.MapGroupOrderDown] utils.HandleRequest", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MGRelation{Id: data.ID})
	if err != nil {
		log.Error("[controller.MapGroupOrderDown] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.MapGroupOrderDownRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.MapGroupOrderDownRequestPartition},
	})

	var model pb.MGRelationMessage

	cn.waitResponse(cc, c, id, &model)

	return
}
