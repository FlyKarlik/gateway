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

func (cn *Controllers) AddGroup(c *gin.Context) {
	log := hclog.Default()

	var data models.Group

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "AddGroup", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controllers.AddGroup] utils.HandleRequest", "error", err)
		return
	}

	dataBytes, err := proto.Marshal(&pb.MGroup{
		Name:         data.Name,
		CreateUserId: "Odil",
		CreateUserIp: c.ClientIP(),
	})
	if err != nil {
		log.Error("[controllers.AddGroup] json.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   dataBytes,
		Key:     []byte{utils.AddGroupRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.AddGroupRequestPartition},
	})

	if err != nil {
		log.Error("[controllers.AddGroup] cn.Producer.SendMessage", "error", err)
		return
	}

	var model pb.GroupMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) Group(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "Group", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.Group] utils.HandleRequest", "error", err)
		return
	}

	groupID := c.Query("id")
	if len(groupID) < 1 {
		log.Error("[controller.Group] c.Query", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MGroup{Id: groupID})
	if err != nil {
		log.Error("[controller.Group] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.GroupRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.GroupRequestPartition},
	})

	var model pb.GroupMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) DeleteGroup(c *gin.Context) {
	var data models.Group
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "DeleteGroup", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.DeleteGroup] utils.HandleRequest", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MGroup{Id: data.ID})
	if err != nil {
		log.Error("[controller.DeleteGroup] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.DeleteGroupRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.DeleteGroupRequestPartition},
	})

	var model pb.GroupMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) Groups(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "Groups", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.Groups] utils.HandleRequest", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   []byte{},
		Key:     []byte{utils.GroupsRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.GroupsRequestPartition},
	})

	var model pb.GroupsMessage
	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) EditGroup(c *gin.Context) {
	var data models.Group
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "EditGroup", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.EditGroup] utils.HandleRequest", "error", err)
		return
	}

	sendBytes, err := proto.Marshal(&pb.MGroup{
		Id:           data.ID,
		Name:         data.Name,
		UpdateUserId: "Odil",
		UpdateUserIp: c.ClientIP(),
	})
	if err != nil {
		log.Error("[controller.EditGroup] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   sendBytes,
		Key:     []byte{utils.EditGroupRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.EditGroupRequestPartition},
	})

	var model pb.GroupMessage
	cn.waitResponse(cc, c, id, &model)

	return
}
