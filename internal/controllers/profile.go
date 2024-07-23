package controllers

import (
	"comet/utils"
	"gateway/internal/controllers/models"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/status"
)

// CheckUsername check username is existed or not
func (cn *Controllers) CheckUsername(c *gin.Context) {
	log := hclog.Default()

	var data models.CheckUsernameRequest
	ctx, span, err := utils.HandleRequest(
		c,
		cn.tracer,
		"CheckUsername",
		&data,
		false,
	)
	defer span.End()
	if err != nil {
		log.Error("[controllers.CheckUsername] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.CheckUsername(ctx, data.Username)
	if err != nil {
		log.Error("[controllers.CheckUsername] cn.client.CheckUsername failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("data", response)
}

// GetAccountInfo get account info
func (cn *Controllers) GetAccountInfo(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(
		c,
		cn.tracer,
		"GetAccountInfo",
		nil,
		true,
	)
	defer span.End()
	if err != nil {
		log.Error("[controllers.GetAccountInfo] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.GetAccountInfo(ctx)
	if err != nil {
		log.Error("[controllers.GetAccountInfo] cn.client.GetAccountInfo failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("data", response)
}

// GetAccountsInfo get account info
func (cn *Controllers) GetAccountsInfo(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(
		c,
		cn.tracer,
		"GetAccountsInfo",
		nil,
		true,
	)
	defer span.End()
	if err != nil {
		log.Error("[controllers.GetAccountsInfo] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.GetAccountsInfo(ctx)
	if err != nil {
		log.Error("[controllers.GetAccountsInfo] cn.client.GetAccountsInfo failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("data", response)
}
