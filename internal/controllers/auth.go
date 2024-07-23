package controllers

import (
	"comet/utils"
	"gateway/internal/controllers/models"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/status"
)

func (cn *Controllers) LoginUser(c *gin.Context) {
	log := hclog.Default()

	var data models.LoginUserRequest
	ctx, span, err := utils.HandleRequest(c, cn.tracer, "LoginUser", &data, false)
	defer span.End()
	if err != nil {
		log.Error("[controllers.LoginUser] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.LoginUser(ctx, data.Email, data.Password)
	if err != nil {
		log.Error("[controllers.LoginUser] cn.client.LoginUser failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("message", response)
}
