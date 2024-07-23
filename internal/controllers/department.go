package controllers

import (
	"comet/utils"
	"gateway/internal/controllers/models"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/status"
)

// AddUserDepartment add new department
func (cn *Controllers) AddUserDepartment(c *gin.Context) {
	log := hclog.Default()

	var data models.AddUserDepartmentRequest
	ctx, span, err := utils.HandleRequest(c, cn.tracer, "AddUserDepartment", &data, true)
	defer span.End()
	if err != nil {
		log.Error("[controllers.AddUserDepartment] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.AddUserDepartment(ctx, data.Name)
	if err != nil {
		log.Error("[controllers.AddUserDepartment] cn.client.AddUserDepartment failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("message", response)
}

// RemoveUserDepartment remove user department by id
func (cn *Controllers) RemoveUserDepartment(c *gin.Context) {
	log := hclog.Default()

	var data models.RemoveUserDepartmentRequest
	ctx, span, err := utils.HandleRequest(c, cn.tracer, "RemoveUserDepartment", &data, true)
	defer span.End()
	if err != nil {
		log.Error("[controllers.RemoveUserDepartment] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.RemoveUserDepartment(ctx, data.Id)
	if err != nil {
		log.Error("[controllers.RemoveUserDepartment] cn.client.RemoveUserDepartment failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("message", response)
}

// GetUserDepartments get all departments for user
func (cn *Controllers) GetUserDepartments(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(c, cn.tracer, "GetUserDepartments", nil, true)
	defer span.End()
	if err != nil {
		log.Error("[controllers.GetUserDepartments] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.GetUserDepartments(ctx)
	if err != nil {
		log.Error("[controllers.GetUserDepartments] cn.client.GetUserDepartments failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("message", response)
}

// GetUserDepartment get user department by ID
func (cn *Controllers) GetUserDepartment(c *gin.Context) {
	log := hclog.Default()

	var data models.GetUserDepartmentRequest
	ctx, span, err := utils.HandleRequest(c, cn.tracer, "GetUserDepartment", &data, true)
	defer span.End()
	if err != nil {
		log.Error("[controllers.GetUserDepartment] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.GetUserDepartment(ctx, data.Id)
	if err != nil {
		log.Error("[controllers.GetUserDepartment] cn.client.GetUserDepartments failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("message", response)
}
