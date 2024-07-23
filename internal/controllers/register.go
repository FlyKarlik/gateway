package controllers

import (
	"comet/utils"
	"gateway/internal/controllers/models"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/status"
)

// RegisterPassword register password for user
func (cn *Controllers) RegisterPassword(c *gin.Context) {
	log := hclog.Default()

	var data models.RegisterPasswordRequest
	ctx, span, err := utils.HandleRequest(
		c,
		cn.tracer,
		"RegisterPassword",
		&data,
		true,
	)
	defer span.End()
	if err != nil {
		log.Error("[controllers.RegisterPassword] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.RegisterPassword(ctx, data.Password)
	if err != nil {
		log.Error("[controllers.RegisterPassword] cn.client.RegisterPassword failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("data", response)
}

// RegisterUsername register username for user
func (cn *Controllers) RegisterUsername(c *gin.Context) {
	log := hclog.Default()

	var data models.RegisterUsernameRequest
	ctx, span, err := utils.HandleRequest(
		c,
		cn.tracer,
		"RegisterUsername",
		&data,
		true,
	)
	defer span.End()
	if err != nil {
		log.Error("[controllers.RegisterUsername] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.RegisterUsername(ctx, data.Username)
	if err != nil {
		log.Error("[controllers.RegisterUsername] cn.client.RegisterUsername failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("data", response)
}

// RegisterUser register user
func (cn *Controllers) RegisterUser(c *gin.Context) {
	log := hclog.Default()

	var data models.RegisterUserRequest
	ctx, span, err := utils.HandleRequest(
		c,
		cn.tracer,
		"RegisterUser",
		&data,
		true,
	)
	defer span.End()
	if err != nil {
		log.Error("[controllers.RegisterUser] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.RegisterUser(
		ctx,
		data.Email,
		data.FirstName,
		data.SecondName,
		data.Password,
		data.DepartmentID,
		data.RoleID,
	)

	if err != nil {
		log.Error("[controllers.RegisterUser] cn.client.RegisterUser failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("data", response)
}

// RemoveUser remove user
func (cn *Controllers) RemoveUser(c *gin.Context) {
	log := hclog.Default()

	var data models.RemoveUserRequest
	ctx, span, err := utils.HandleRequest(c, cn.tracer, "RemoveUser", &data, true)
	defer span.End()
	if err != nil {
		log.Error("[controllers.RemoveUser] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.RemoveUser(ctx, data.Id)
	if err != nil {
		log.Error("[controllers.RemoveUser] cn.client.RemoveUser failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("data", response)
}
