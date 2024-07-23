package controllers

import (
	"bytes"
	"comet/utils"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gateway/internal/client"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
	"image"
	"image/draw"
	"image/png"
	pb "protos/maps"
	"strings"
)

func (cn *Controllers) Pattern(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "Pattern", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.Pattern] utils.HandleRequest", "error", err)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	mode := c.Query("mode")
	switch mode {
	case "":
		err = cn.Producer.SendMessage(&kafka.Message{
			Value:   []byte{},
			Key:     []byte{utils.PatternsRequest},
			Headers: []kafka.Header{{Key: id}},
			TopicPartition: kafka.TopicPartition{
				Topic:     &cn.cfg.KafkaMapsRequestTopic,
				Partition: utils.PatternsRequestPartition},
		})

		if err != nil {
			log.Error("[controllers.Pattern] cn.Producer.SendMessage", "error", err)
			return
		}

		var model pb.PatternsMessage

		cn.waitResponse(cc, c, id, &model)

		return
	case "single":
		patternID := c.Query("id")

		if len(patternID) < 1 {
			log.Error("[controllers.Pattern] strconv.Atoi", "error", err.Error())
			c.Set("message", "strconv.Atoi conversion failed")
			c.Set("code", utils.CodeInvalidArgument)
			c.Set("status", "failed")

			break
		}

		dataBytes, err := proto.Marshal(&pb.Pattern{Id: patternID})
		if err != nil {
			log.Error("[controller.Pattern] proto.Marshal", "error", err)
			return
		}

		err = cn.Producer.SendMessage(&kafka.Message{
			Value:   dataBytes,
			Key:     []byte{utils.PatternRequest},
			Headers: []kafka.Header{{Key: id}},
			TopicPartition: kafka.TopicPartition{
				Topic:     &cn.cfg.KafkaMapsRequestTopic,
				Partition: utils.PatternRequestPartition},
		})

		if err != nil {
			log.Error("[controllers.Pattern] cn.Producer.SendMessage", "error", err)
			return
		}

		var model pb.PatternMessage

		cn.waitResponse(cc, c, id, &model)

		return
	}
}

func (cn *Controllers) Sprite(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "Sprite", nil, false)
	defer span.End()
	if err != nil {
		log.Error("[controller.Sprite] utils.HandleRequest", "error", err)
		c.Set("code", utils.CodeInvalidArgument)
		return
	}

	id := uuid.NewString()
	cc, cancel := context.WithTimeout(ctx, client.DialTimeout)
	defer cancel()

	err = cn.Producer.SendMessage(&kafka.Message{
		Value:   []byte{},
		Key:     []byte{utils.PatternsRequest},
		Headers: []kafka.Header{{Key: id}},
		TopicPartition: kafka.TopicPartition{
			Topic:     &cn.cfg.KafkaMapsRequestTopic,
			Partition: utils.PatternsRequestPartition},
	})

	if err != nil {
		log.Error("[controllers.Sprite] cn.Producer.SendMessage", "error", err)
		return
	}

	var model pb.PatternsMessage

	cn.waitResponse(cc, c, id, &model)

	if len(model.Patterns) < 1 {
		log.Error("[controllers.Sprite] len(model.Patterns) less then 1", "error", err)
		c.Set("message", "decode image failed")
		c.Set("code", utils.CodeInvalidArgument)
		c.Set("status", "failed")
		c.Set("data", nil)
		return
	}

	data := "{"

	sprite := image.NewRGBA(image.Rectangle{
		Min: image.Point{},
		Max: image.Point{},
	})

	var size int

	for i := 0; i < len(model.Patterns); i++ {
		img, err := base64.StdEncoding.DecodeString(model.Patterns[i].Img)

		if err != nil {
			log.Error("[controllers.Sprite] base64.StdEncoding.DecodeString", "error", err)
			c.Set("message", "decode image failed")
			c.Set("code", utils.CodeInvalidArgument)
			c.Set("status", "failed")
			c.Set("data", nil)
			return
		}

		r := bytes.NewReader(img)
		png2, err := png.Decode(r)

		sp2 := image.Point{X: sprite.Bounds().Dx()}

		r2 := image.Rectangle{Min: sp2, Max: sp2.Add(png2.Bounds().Size())}

		re := image.Rectangle{Min: image.Point{}, Max: image.Point{X: r2.Max.X, Y: 200}}

		rgba := image.NewRGBA(re)

		draw.Draw(rgba, sprite.Bounds(), sprite, image.Point{}, draw.Src)
		draw.Draw(rgba, r2, png2, image.Point{}, draw.Src)

		sprite = rgba

		data = data + fmt.Sprintf(`"%s":{ "name": "%s", "width":%v, "height":%v, "pixelRatio":2, "visible": true, "x":%v, "y":%v}`,
			model.Patterns[i].Id,
			model.Patterns[i].Name,
			model.Patterns[i].X,
			model.Patterns[i].Y,
			size,
			0,
		)
		if i == len(model.Patterns)-1 {
			data = data + "}"
		} else {
			data = data + ","
		}

		size = size + int(model.Patterns[i].X)
	}

	buffer := new(bytes.Buffer)
	err = png.Encode(buffer, sprite)

	if err != nil {
		log.Error("[controllers.Sprite] png.Encode", "error", err)
		c.Set("code", utils.CodeInternal)
		return
	}

	dataBytes := make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &dataBytes)
	if err != nil {
		log.Error("[controllers.Sprite] json.Unmarshal", "error", err)
		c.Set("message", "unmarshal ")
		c.Set("code", utils.CodeInvalidArgument)
		c.Set("status", "failed")
		c.Set("data", nil)
		return
	}

	switch {
	case strings.Contains(c.Request.URL.String(), "json"):
		c.Set("lock", true)
		c.Header("Content-Type", "application/json")
		c.JSON(utils.ConvertToHttpCode(utils.CodeOK), dataBytes)
	case strings.Contains(c.Request.URL.String(), "png"):
		c.Set("lock", true)
		c.Header("Content-Type", "image/png")
		c.Data(utils.ConvertToHttpCode(utils.CodeOK), "image/png", buffer.Bytes())
	default:
		c.Set("lock", true)
		c.Header("Content-Type", "application/json")
		c.JSON(utils.ConvertToHttpCode(utils.CodeOK), dataBytes)
	}
}
