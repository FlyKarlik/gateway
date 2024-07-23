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
	"strconv"
	"time"
)

func (cn *Controllers) Tables(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "Tables", nil, false)
	defer span.End()

	if err != nil {
		log.Error("[controller.Tables] utils.HandleRequest failed", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   []byte{},
		Key:     []byte{utils.TablesRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.TablesRequestPartition},
	})

	if err != nil {
		log.Error("[controllers.Tables] cn.Producer.SendMessage", "error", err)
		return
	}

	var model pb.TablesMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) Table(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "Table", nil, false)
	defer span.End()

	if err != nil {
		log.Error("[controller.Table] utils.HandleRequest failed", "error", err)
		return
	}

	tableID := c.Query("id")
	if len(tableID) < 1 {
		log.Error("[controller.Table] c.Query failed", "error", "bad request")
		return
	}

	dataBytes, err := proto.Marshal(&pb.Table{Id: tableID})
	if err != nil {
		log.Error("[controllers.Table] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   dataBytes,
		Key:     []byte{utils.TableRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.TableRequestPartition},
	})

	if err != nil {
		log.Error("[controller.Table] cn.Producer.SendMessage failed", "error", err)
		return
	}

	var model pb.TableMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) DeleteTable(c *gin.Context) {
	log := hclog.Default()

	var table pb.Table

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "DeleteTable", &table, false)
	defer span.End()

	if err != nil {
		log.Error("[controller.DeleteTable] utils.HandleRequest failed", "error", err)
		return
	}

	dataBytes, err := proto.Marshal(&table)
	if err != nil {
		log.Error("[controllers.DeleteTable] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   dataBytes,
		Key:     []byte{utils.DeleteTableRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.DeleteTableRequestPartition},
	})

	if err != nil {
		log.Error("[controller.DeleteTable] cn.Producer.SendMessage failed", "error", err)
		return
	}

	var model pb.TableMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) AddTable(c *gin.Context) {
	log := hclog.Default()

	var table models.Table

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "AddTable", &table, false)
	defer span.End()

	if err != nil {
		log.Error("[controller.AddTable] utils.HandleRequest failed", "error", err)
		return
	}

	dataBytes, err := proto.Marshal(&pb.Table{
		Name:               table.Name,
		Alias:              table.Alias,
		IsRelated:          table.IsRelated,
		IsVersioned:        table.IsVersioned,
		IsArchived:         table.IsArchived,
		IsGeometryNullable: table.IsGeometryNullable,
		GeometryType:       table.GeometryType,
		Srid:               table.SRID,
		TableType:          table.TableType,
	})
	if err != nil {
		log.Error("[controllers.AddTable] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   dataBytes,
		Key:     []byte{utils.AddTableRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.AddTableRequestPartition},
	})

	if err != nil {
		log.Error("[controller.AddTable] cn.Producer.SendMessage failed", "error", err)
		return
	}

	var model pb.TableMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) EditTable(c *gin.Context) {
	log := hclog.Default()

	var table models.Table

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "EditTable", &table, false)
	defer span.End()

	if err != nil {
		log.Error("[controller.EditTable] utils.HandleRequest failed", "error", err)
		return
	}

	dataBytes, err := proto.Marshal(&pb.Table{
		Id:                 table.ID,
		Name:               table.Name,
		Alias:              table.Alias,
		IsRelated:          table.IsRelated,
		IsVersioned:        table.IsVersioned,
		IsArchived:         table.IsArchived,
		IsGeometryNullable: table.IsGeometryNullable,
		GeometryType:       table.GeometryType,
		Srid:               table.SRID,
		TableType:          table.TableType,
	})
	if err != nil {
		log.Error("[controllers.EditTable] proto.Marshal", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   dataBytes,
		Key:     []byte{utils.EditTableRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.EditTableRequestPartition},
	})

	if err != nil {
		log.Error("[controller.EditTable] cn.Producer.SendMessage failed", "error", err)
		return
	}

	var model pb.TableMessage

	cn.waitResponse(cc, c, id, &model)

	return

}

func (cn *Controllers) TableColumns(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "TableColumns", nil, false)
	defer span.End()

	if err != nil {
		log.Error("[controller.TableColumns] utils.HandleRequest failed", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	tableID := c.Query("table")
	if len(tableID) < 1 {
		log.Error("[controller.TableColumns] c.Query failed", "error", "bad request")
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	dataBytes, err := proto.Marshal(&pb.Table{Id: tableID})
	if err != nil {
		log.Error("[controllers.TableColumns] proto.Marshal", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   dataBytes,
		Key:     []byte{utils.TableColumnsRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.TableColumnsRequestPartition},
	})

	if err != nil {
		log.Error("[controller.TableColumns] cn.Producer.SendMessage failed", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	var model pb.ColumnsMessage

	cn.waitResponse(cc, c, id, &model)

	return
}

func (cn *Controllers) TableColumnUniqueValues(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "TableColumnUniqueValues", nil, false)
	defer span.End()

	if err != nil {
		log.Error("[controller.TableColumnUniqueValues] utils.HandleRequest failed", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	tableName := c.Query("table_name")
	if len(tableName) < 1 {
		log.Error("[controller.TableColumnUniqueValues] len(tableID) < 1 failed")
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	columnName := c.Query("column_name")
	if len(columnName) < 1 {
		log.Error("[controller.TableColumnUniqueValues] len(columnName) < 1 failed")
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	dataBytes, err := proto.Marshal(&pb.ColumnUnique{TableName: tableName, ColumnName: columnName})
	if err != nil {
		log.Error("[controllers.TableColumnUniqueValues] proto.Marshal", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   dataBytes,
		Key:     []byte{utils.TableColumnUniqueValuesRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.TableColumnUniqueValuesRequestPartition},
	})

	if err != nil {
		log.Error("[controller.TableColumnUniqueValues] cn.Producer.SendMessage failed", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	var model pb.ColumnUniqueMessage

	cn.waitResponse(cc, c, id, &model)

	return

}

func (cn *Controllers) GetFeatures(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "GetFeatures", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.GetFeatures] utils.HandleRequest failed", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	layers, ok := c.GetQuery("layers")
	if ok != true {
		log.Error("[controller.GetFeatures] tables not send")
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	xMin, ok := c.GetQuery("xmin")
	if ok != true {
		log.Error("[controller.GetFeatures] xmin not send")
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	xMax, ok := c.GetQuery("xmax")
	if ok != true {
		log.Error("[controller.GetFeatures] xmax not send")
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	yMin, ok := c.GetQuery("ymin")
	if ok != true {
		log.Error("[controller.GetFeatures] ymin not send")
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	yMax, ok := c.GetQuery("ymax")
	if ok != true {
		log.Error("[controller.GetFeatures] ymax not send")
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	var params []float64

	t, err := strconv.ParseFloat(xMin, 64)
	if err != nil {
		log.Error("[controller.GetFeatures] strconv.ParseFloat failed")
		c.Set("code", utils.CodeInvalidArgument)
		return
	} else {
		params = append(params, t)
	}

	t, err = strconv.ParseFloat(xMax, 64)
	if err != nil {
		log.Error("[controller.GetFeatures] strconv.ParseFloat failed")
		c.Set("code", utils.CodeInvalidArgument)
		return
	} else {
		params = append(params, t)
	}

	t, err = strconv.ParseFloat(yMin, 64)
	if err != nil {
		log.Error("[controller.GetFeatures] strconv.ParseFloat failed")
		c.Set("code", utils.CodeInvalidArgument)
		return
	} else {
		params = append(params, t)
	}

	t, err = strconv.ParseFloat(yMax, 64)
	if err != nil {
		log.Error("[controller.GetFeatures] strconv.ParseFloat failed")
		c.Set("code", utils.CodeInvalidArgument)
		return
	} else {
		params = append(params, t)
	}

	dataBytes, err := proto.Marshal(&pb.TableFeaturesRequest{
		Layers: layers,
		Xmin:   float32(params[0]),
		Xmax:   float32(params[1]),
		Ymin:   float32(params[2]),
		Ymax:   float32(params[3]),
	})
	if err != nil {
		log.Error("[controllers.GetFeatures] proto.Marshal", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   dataBytes,
		Key:     []byte{utils.TableFeaturesRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.TableFeaturesRequestPartition},
	})

	if err != nil {
		log.Error("[controller.GetFeatures] cn.Producer.SendMessage failed", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	var model pb.TableFeatureMessage

	cn.waitResponse(cc, c, id, &model)

	return

}

func (cn *Controllers) TableSwitcher(c *gin.Context, n client.Node, model interface{}) bool {
	switch {
	case n.P == utils.AddTableResponsePartition || n.P == utils.TableResponsePartition ||
		n.P == utils.EditTableResponsePartition || n.P == utils.DeleteTableResponsePartition:
		var m pb.TableMessage
		err := proto.Unmarshal(n.Message, &m)
		if err != nil {
			c.Set("code", utils.CodeInternal)
			return false
		}

		c.Set("code", m.Code)
		c.Set("data", m.Table)
		return false

	case n.P == utils.TablesResponsePartition:
		var m pb.TablesMessage
		err := proto.Unmarshal(n.Message, &m)
		if err != nil {
			c.Set("code", utils.CodeInternal)
			return false
		}

		c.Set("code", m.Code)
		c.Set("data", m.Tables)
		return false

	case n.P == utils.TableColumnsResponsePartition:
		err := proto.Unmarshal(n.Message, model.(*pb.ColumnsMessage))
		if err != nil {
			c.Set("code", utils.CodeInternal)
			return false
		}

		c.Set("code", model.(*pb.ColumnsMessage).Code)
		c.Set("data", model.(*pb.ColumnsMessage).Columns)
		return false

	case n.P == utils.TableColumnUniqueValuesResponsePartition:
		err := proto.Unmarshal(n.Message, model.(*pb.ColumnUniqueMessage))
		if err != nil {
			c.Set("code", utils.CodeInternal)
			return false
		}

		c.Set("code", model.(*pb.ColumnUniqueMessage).Code)
		c.Set("data", model.(*pb.ColumnUniqueMessage).Unique)
		return false

	case n.P == utils.TableFeaturesResponsePartition:
		err := proto.Unmarshal(n.Message, model.(*pb.TableFeatureMessage))
		if err != nil {
			c.Set("code", utils.CodeInternal)
			return false
		}

		c.Set("code", model.(*pb.TableFeatureMessage).Code)
		c.Set("data", model.(*pb.TableFeatureMessage).Features)
		return false

	default:
		return true
	}
}
