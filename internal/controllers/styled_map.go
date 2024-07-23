package controllers

import (
	"comet/utils"
	"context"
	"gateway/internal/client"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	pb "protos/maps"
)

func (cn *Controllers) StyledMap(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "StyledMap", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.StyledMap] utils.HandleRequest", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	mapID := c.Query("id")
	if len(mapID) < 1 {
		log.Error("[controller.StyledMap] c.Query(id)", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MMap{Id: mapID})
	if err != nil {
		log.Error("[controller.StyledMap] proto.Marshal", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.StyledMapRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.StyledMapRequestPartition},
	})

	var model pb.StyledMapMessage
	cn.waitResponse(cc, c, id, &model)

	return
}
