package controllers

import (
	"comet/utils"
	"gateway/internal/controllers/models"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/status"
)

// RefreshToken refresh auth or access token
func (cn *Controllers) RefreshToken(c *gin.Context) {
	log := hclog.Default()

	var data models.RefreshTokenRequest
	ctx, span, err := utils.HandleRequest(
		c,
		cn.tracer,
		"RefreshToken",
		&data,
		false,
	)
	defer span.End()
	if err != nil {
		log.Error("[controllers.RefreshToken] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.RefreshToken(ctx, data.RefreshToken)
	if err != nil {
		log.Error("[controllers.RefreshToken] cn.client.RefreshToken failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("data", response)
}
