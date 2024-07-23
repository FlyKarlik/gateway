package controllers

import (
	"comet/utils"
	"gateway/internal/controllers/models"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/status"
)

// AddUserRole add new role
func (cn *Controllers) AddUserRole(c *gin.Context) {
	log := hclog.Default()

	var data models.AddUserRoleRequest
	ctx, span, err := utils.HandleRequest(c, cn.tracer, "AddUserRole", &data, true)
	defer span.End()
	if err != nil {
		log.Error("[controllers.AddUserRole] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.AddUserRole(ctx, data.Name)
	if err != nil {
		log.Error("[controllers.AddUserRole] cn.client.AddUserRole failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("message", response)
}

// RemoveUserRole remove user role by id
func (cn *Controllers) RemoveUserRole(c *gin.Context) {
	log := hclog.Default()

	var data models.RemoveUserRoleRequest
	ctx, span, err := utils.HandleRequest(c, cn.tracer, "RemoveUserRole", &data, true)
	defer span.End()
	if err != nil {
		log.Error("[controllers.RemoveUserRole] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.RemoveUserRole(ctx, data.Id)
	if err != nil {
		log.Error("[controllers.RemoveUserRole] cn.client.RemoveUserRole failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("message", response)
}

// GetUserRoles get all roles for user
func (cn *Controllers) GetUserRoles(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "GetUserRoles", nil, true)
	defer span.End()
	if err != nil {
		log.Error("[controllers.GetUserRoles] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.GetUserRoles(ctx)
	if err != nil {
		log.Error("[controllers.GetUserRoles] cn.client.GetUserRoles failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("message", response)
}

// GetUserRole remove user role by id
func (cn *Controllers) GetUserRole(c *gin.Context) {
	log := hclog.Default()

	var data models.GetUserRoleRequest
	ctx, span, err := utils.HandleRequest(c, cn.tracer, "GetUserRole", &data, true)
	defer span.End()
	if err != nil {
		log.Error("[controllers.GetUserRole] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.GetUserRole(ctx, data.Id)
	if err != nil {
		log.Error("[controllers.GetUserRole] cn.client.GetUserRole failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("message", response)
}
