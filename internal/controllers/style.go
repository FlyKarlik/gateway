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

func (cn *Controllers) AddStyle(c *gin.Context) {
	log := hclog.Default()

	var data models.Style

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "AddStyle", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controllers.AddStyle] utils.HandleRequest", "error", err)
		return
	}

	dataBytes, err := proto.Marshal(data.ToMStyle())
	if err != nil {
		log.Error("[controllers.AddStyle] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   dataBytes,
		Key:     []byte{utils.AddStyleRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.AddStyleRequestPartition},
	})

	if err != nil {
		log.Error("[controllers.AddStyle] cn.Producer.SendMessage", "error", err)
		return
	}

	var model pb.StyleMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) Style(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "Style", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.Style] utils.HandleRequest", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	styleID := c.Query("id")
	if len(styleID) < 1 {
		log.Error("[controller.Style] passed id is bad")
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MStyle{Id: styleID})
	if err != nil {
		log.Error("[controller.Style] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.StyleRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.StyleRequestPartition},
	})

	var model pb.StyleMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) DeleteStyle(c *gin.Context) {
	var data models.Style
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "DeleteStyle", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.DeleteStyle] utils.HandleRequest", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MStyle{Id: data.ID})
	if err != nil {
		log.Error("[controller.DeleteStyle] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.DeleteStyleRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.DeleteStyleRequestPartition},
	})

	var model pb.StyleMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) Styles(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "Style", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.Styles] utils.HandleRequest", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   []byte{},
		Key:     []byte{utils.StylesRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.StylesRequestPartition},
	})

	var model pb.StylesMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) StylesPagination(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "Style", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.StylesPagination] utils.HandleRequest", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	pageSize := c.Query("page_size")
	page := c.Query("page")

	if len(pageSize) < 1 {
		log.Error("[controller.StylesPagination] page size is zero or nil", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	if len(page) < 1 {
		log.Error("[controller.StylesPagination] page is nil", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	sendBytes, err := proto.Marshal(&pb.StylesPagination{
		PageSize: pageSize,
		Page:     page,
	})
	if err != nil {
		log.Error("[controller.StylesPagination] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.StylesPaginationRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.StylesPaginationRequestPartition},
	})

	var model pb.StylesPaginationMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) EditStyle(c *gin.Context) {
	var data models.Style
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "EditStyle", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.EditStyle] utils.HandleRequest", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(data.ToMStyle())
	if err != nil {
		log.Error("[controller.EditStyle] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.EditStyleRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.EditStyleRequestPartition},
	})

	var model pb.StyleMessage
	cn.waitResponse(cc, c, id, &model)

	return
}
