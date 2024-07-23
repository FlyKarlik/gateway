package controllers

import (
	"comet/utils"
	"context"
	"gateway/config"
	"gateway/internal/client"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"go.opentelemetry.io/otel/trace"
	pb "protos/maps"
	"strings"
)

// Controllers for account gateway
type Controllers struct {
	tracer       trace.Tracer
	cfg          *config.Config
	Producer     *client.ServiceProducer
	responseHash *client.MessageHash
	client       *client.AccountClient
}

func NewControllers(tracer trace.Tracer, cfg *config.Config, p *client.ServiceProducer, responseHash *client.MessageHash, client *client.AccountClient) *Controllers {
	return &Controllers{
		tracer:       tracer,
		Producer:     p,
		responseHash: responseHash,
		cfg:          cfg,
		client:       client,
	}
}

func (cn *Controllers) waitResponse(ctx context.Context, c *gin.Context, id string, model interface{}) {
	var run = true
	if model != nil {
		for run {
			select {
			case <-ctx.Done():
				c.Set("code", utils.CodeDeadlineExceeded)
				run = false
				break
			default:
				if v, ok := cn.responseHash.Get(id); ok == true {
					n := v.Value.(client.Node)
					switch {
					case n.P == utils.AddLayerResponsePartition || n.P == utils.LayerResponsePartition ||
						n.P == utils.EditLayerResponsePartition || n.P == utils.DeleteLayerResponsePartition:
						var layer pb.LayerMessage
						err := proto.Unmarshal(n.Message, &layer)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", layer.Code)
						c.Set("data", layer.Layer)
						run = false
						break

					case n.P == utils.LayersResponsePartition:
						var layer pb.LayersMessage
						err := proto.Unmarshal(n.Message, &layer)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", layer.Code)
						c.Set("data", layer.Layers)
						run = false
						break

					case n.P == utils.AddGroupResponsePartition || n.P == utils.GroupResponsePartition ||
						n.P == utils.EditGroupResponsePartition || n.P == utils.DeleteGroupResponsePartition:
						var group pb.GroupMessage
						err := proto.Unmarshal(n.Message, &group)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", group.Code)
						c.Set("data", group.Group)
						run = false
						break

					case n.P == utils.GroupsResponsePartition:
						var group pb.GroupsMessage
						err := proto.Unmarshal(n.Message, &group)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", group.Code)
						c.Set("data", group.Groups)
						run = false
						break

					case n.P == utils.AddStyleResponsePartition || n.P == utils.StyleResponsePartition ||
						n.P == utils.EditStyleResponsePartition || n.P == utils.DeleteStyleResponsePartition:
						var style pb.StyleMessage
						err := proto.Unmarshal(n.Message, &style)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", style.Code)
						c.Set("data", style.Style)
						run = false
						break

					case n.P == utils.StylesResponsePartition:
						var styles pb.StylesMessage
						err := proto.Unmarshal(n.Message, &styles)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", styles.Code)
						c.Set("data", styles.Styles)
						run = false
						break

					case n.P == utils.StylesPaginationResponsePartition:
						err := proto.Unmarshal(n.Message, model.(*pb.StylesPaginationMessage))
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", model.(*pb.StylesPaginationMessage).Code)
						c.Set("data", model.(*pb.StylesPaginationMessage).Styles)
						run = false
						break

					case n.P == utils.AddMapResponsePartition || n.P == utils.MapResponsePartition ||
						n.P == utils.EditMapResponsePartition || n.P == utils.DeleteMapResponsePartition:
						var mmap pb.MapMessage
						err := proto.Unmarshal(n.Message, &mmap)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", mmap.Code)
						c.Set("data", mmap.Map)
						run = false
						break

					case n.P == utils.MapsResponsePartition:
						var Maps pb.MapsMessage
						err := proto.Unmarshal(n.Message, &Maps)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", Maps.Code)
						c.Set("data", Maps.Maps)
						run = false
						break

					case n.P == utils.AddGroupLayerRelationResponsePartition || n.P == utils.DeleteGroupLayerRelationResponsePartition:
						var m pb.GroupLayerRelationMessage
						err := proto.Unmarshal(n.Message, &m)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", m.Code)
						c.Set("data", m.Relation)
						run = false
						break

					case n.P == utils.GroupLayerRelationsResponsePartition:
						var m pb.GroupLayerRelationsMessage
						err := proto.Unmarshal(n.Message, &m)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", m.Code)
						c.Set("data", m.Relations)
						run = false
						break

					case n.P == utils.LayerRelationGroupsResponsePartition:
						var m pb.GroupsMessage
						err := proto.Unmarshal(n.Message, &m)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", m.Code)
						c.Set("data", m.Groups)
						run = false
						break

					case n.P == utils.GroupRelationLayersResponsePartition:
						var m pb.LayersMessage
						err := proto.Unmarshal(n.Message, &m)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", m.Code)
						c.Set("data", m.Layers)
						run = false
						break

					case n.P == utils.AddMapGroupRelationResponsePartition || n.P == utils.DeleteMapGroupRelationResponsePartition:
						var m pb.MGRelationMessage
						err := proto.Unmarshal(n.Message, &m)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", m.Code)
						c.Set("data", m.Relation)
						run = false
						break

					case n.P == utils.MapGroupRelationsResponsePartition:
						var m pb.MGRelationsMessage
						err := proto.Unmarshal(n.Message, &m)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", m.Code)
						c.Set("data", m.Relations)
						run = false
						break

					case n.P == utils.MapRelationGroupsResponsePartition:
						var m pb.GroupsMessage
						err := proto.Unmarshal(n.Message, &m)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", m.Code)
						c.Set("data", m.Groups)
						run = false
						break

					case n.P == utils.GroupRelationMapsResponsePartition:
						var m pb.MapsMessage
						err := proto.Unmarshal(n.Message, &m)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", m.Code)
						c.Set("data", m.Maps)
						run = false
						break

					case n.P == utils.MapGroupOrderUpResponsePartition:
						err := proto.Unmarshal(n.Message, model.(*pb.MGRelationMessage))
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", model.(*pb.MGRelationMessage).Code)
						c.Set("data", model.(*pb.MGRelationMessage).Relation)
						run = false
						break

					case n.P == utils.MapGroupOrderDownResponsePartition:
						err := proto.Unmarshal(n.Message, model.(*pb.MGRelationMessage))
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", model.(*pb.MGRelationMessage).Code)
						c.Set("data", model.(*pb.MGRelationMessage).Relation)
						run = false
						break

					case n.P == utils.AddLayerStyleRelationResponsePartition || n.P == utils.DeleteLayerStyleRelationResponsePartition:
						var m pb.LSRMessage
						err := proto.Unmarshal(n.Message, &m)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", m.Code)
						c.Set("data", m.Relation)
						run = false
						break

					case n.P == utils.LayerStyleRelationsResponsePartition:
						var m pb.LSRsMessage
						err := proto.Unmarshal(n.Message, &m)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", m.Code)
						c.Set("data", m.Relations)
						run = false
						break

					case n.P == utils.LayerRelationStylesResponsePartition:
						var m pb.StylesMessage
						err := proto.Unmarshal(n.Message, &m)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", m.Code)
						c.Set("data", m.Styles)
						run = false
						break

					case n.P == utils.StyleRelationLayersResponsePartition:
						var m pb.LayersMessage
						err := proto.Unmarshal(n.Message, &m)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", m.Code)
						c.Set("data", m.Layers)
						run = false
						break

					case n.P == utils.StyledMapResponsePartition:
						var m pb.StyledMapMessage
						err := proto.Unmarshal(n.Message, &m)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", m.Code)
						c.Set("data", m.Map)
						run = false
						break

					case n.P == utils.AddPatternResponsePartition || n.P == utils.PatternResponsePartition || n.P == utils.DeletePatternResponsePartition:
						var m pb.PatternMessage
						err := proto.Unmarshal(n.Message, &m)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						c.Set("code", m.Code)
						c.Set("data", m.Pattern)
						run = false
						break

					case n.P == utils.PatternsResponsePartition:
						var m = model.(*pb.PatternsMessage)
						err := proto.Unmarshal(n.Message, m)
						if err != nil {
							c.Set("code", utils.CodeInternal)
							run = false
							break
						}

						if !strings.Contains(c.Request.URL.String(), "sprite") {
							c.Set("code", m.Code)
							c.Set("data", m.Patterns)
						}

						run = false
						break
					}

					if ok := cn.TableSwitcher(c, n, model); ok == true {
						run = false
						break
					}

					go func() {
						cn.responseHash.Delete(id)
					}()
					run = false
					break
				}
			}
		}
	}
	return
}
