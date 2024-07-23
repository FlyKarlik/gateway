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

func (cn *Controllers) LayerStyleRelations(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "LayerStyleRelations", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.LayerStyleRelations] utils.HandleRequest", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   []byte{},
		Key:     []byte{utils.LayerStyleRelationsRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.LayerStyleRelationsRequestPartition},
	})

	var model pb.LSRsMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) AddLayerStyleRelation(c *gin.Context) {
	log := hclog.Default()

	var data models.LayerStyleRelation

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "AddLayerStyleRelation", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controllers.AddLayerStyleRelation] utils.HandleRequest", "error", err)
		return
	}

	dataBytes, err := proto.Marshal(&pb.LSRelation{
		LayerId: data.LayerID,
		StyleId: data.StyleID,
	})
	if err != nil {
		log.Error("[controllers.AddLayerStyleRelation] json.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   dataBytes,
		Key:     []byte{utils.AddLayerStyleRelationRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.AddLayerStyleRelationRequestPartition},
	})

	if err != nil {
		log.Error("[controllers.AddLayerStyleRelation] cn.Producer.SendMessage", "error", err)
		return
	}

	var model pb.LSRMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) DeleteLayerStyleRelation(c *gin.Context) {
	var data models.LayerStyleRelation
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "DeleteLayerStyleRelation", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.DeleteLayerStyleRelation] utils.HandleRequest", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.LSRelation{Id: data.ID})
	if err != nil {
		log.Error("[controller.DeleteLayerStyleRelation] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.DeleteLayerStyleRelationRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.DeleteLayerStyleRelationRequestPartition},
	})

	var model pb.LSRMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) LayerRelationStyles(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "LayerRelationStyles", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.LayerRelationStyles] utils.HandleRequest", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	layerID := c.Query("id")
	if len(layerID) < 1 {
		log.Error("[controller.LayerRelationStyles] c.Query(id)", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MLayer{Id: layerID})
	if err != nil {
		log.Error("[controller.LayerRelationStyles] proto.Marshal", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.LayerRelationStylesRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.LayerRelationStylesRequestPartition},
	})

	var model pb.LRSMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) StyleRelationLayers(c *gin.Context) {
	log := hclog.Default()
	var data pb.MStyle

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "StyleRelationLayers", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.StyleRelationLayers] utils.HandleRequest", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&data)
	if err != nil {
		log.Error("[controller.StyleRelationLayers] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.StyleRelationLayersRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.StyleRelationLayersRequestPartition},
	})

	var model pb.SRLMessage
	cn.waitResponse(cc, c, id, &model)

	return
}
