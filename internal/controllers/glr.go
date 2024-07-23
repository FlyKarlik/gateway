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

func (cn *Controllers) AddGroupLayerRelation(c *gin.Context) {
	log := hclog.Default()

	var data models.GroupLayerRelation

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "AddGroupLayerRelation", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controllers.AddGroupLayerRelation] utils.HandleRequest", "error", err)
		return
	}

	dataBytes, err := proto.Marshal(&pb.GroupLayerRelation{
		GroupId: *data.GroupID,
		LayerId: *data.LayerID,
	})
	if err != nil {
		log.Error("[controllers.AddGroupLayerRelation] json.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   dataBytes,
		Key:     []byte{utils.AddGroupLayerRelationRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.AddGroupLayerRelationRequestPartition},
	})

	if err != nil {
		log.Error("[controllers.AddGroupLayerRelation] cn.Producer.SendMessage", "error", err)
		return
	}

	var model pb.GroupLayerRelationMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) GroupLayerRelations(c *gin.Context) {

	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "GroupLayerRelation", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.GroupLayerRelation] utils.HandleRequest", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   []byte{},
		Key:     []byte{utils.GroupLayerRelationsRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.GroupLayerRelationsRequestPartition},
	})

	var model pb.GroupLayerRelationsMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) DeleteGroupLayerRelation(c *gin.Context) {
	var data models.GroupLayerRelation
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "DeleteGroupLayerRelation", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.DeleteGroupLayerRelation] utils.HandleRequest", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.GroupLayerRelation{GroupId: *data.GroupID, LayerId: *data.LayerID})
	if err != nil {
		log.Error("[controller.DeleteGroupLayerRelation] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.DeleteGroupLayerRelationRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.DeleteGroupLayerRelationRequestPartition},
	})

	var model pb.GroupLayerRelationMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) GroupRelationLayers(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "GroupRelationLayers", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.GroupRelationLayers] utils.HandleRequest", "error", err)
		return
	}

	groupID := c.Query("id")
	if len(groupID) < 1 {
		log.Error("[controller.GroupRelationLayers] c.Query", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MGroup{Id: groupID})
	if err != nil {
		log.Error("[controller.GroupRelationLayers] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.GroupRelationLayersRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.GroupRelationLayersRequestPartition},
	})

	var model pb.GLMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) LayerRelationGroups(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "LayerRelationGroups", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.LayerRelationGroups] utils.HandleRequest", "error", err)
		return
	}

	layerID := c.Query("id")
	if len(layerID) < 1 {
		log.Error("[controller.LayerRelationGroups] c.Query", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MLayer{Id: layerID})
	if err != nil {
		log.Error("[controller.LayerRelationGroups] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.LayerRelationGroupsRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.LayerRelationGroupsRequestPartition},
	})

	var model pb.LGMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) GroupLayerOrderUp(c *gin.Context) {
	var data models.GroupLayerRelation
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "GroupLayerOrderUp", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.GroupLayerOrderUp] utils.HandleRequest", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.GroupLayerRelation{Id: *data.ID})
	if err != nil {
		log.Error("[controller.GroupLayerOrderUp] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.GroupLayerOrderUpRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.GroupLayerOrderUpRequestPartition},
	})

	var model pb.GroupLayerRelationMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) GroupLayerOrderDown(c *gin.Context) {
	var data models.GroupLayerRelation
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "GroupLayerOrderDown", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.GroupLayerOrderDown] utils.HandleRequest", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.GroupLayerRelation{Id: *data.ID})
	if err != nil {
		log.Error("[controller.GroupLayerOrderDown] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.GroupLayerOrderDownRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.GroupLayerOrderDownRequestPartition},
	})

	var model pb.GroupLayerRelationMessage

	cn.waitResponse(cc, c, id, &model)

	return
}
