package controllers

import (
	"comet/utils"
	"gateway/internal/controllers/models"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/status"
)

// ForgotPassword get forgot password token
func (cn *Controllers) ForgotPassword(c *gin.Context) {
	log := hclog.Default()

	ctx, span, err := utils.HandleRequest(
		c,
		cn.tracer,
		"ForgotPassword",
		nil,
		true,
	)
	defer span.End()
	if err != nil {
		log.Error("[controllers.ForgotPassword] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.ForgotPassword(ctx)
	if err != nil {
		log.Error("[controllers.ForgotPassword] cn.client.ForgotPassword failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("data", response)
}

// VerifyForgotPassword verify forgot process
func (cn *Controllers) VerifyForgotPassword(c *gin.Context) {
	log := hclog.Default()

	var data models.VerifyForgotPasswordRequest
	ctx, span, err := utils.HandleRequest(
		c,
		cn.tracer,
		"VerifyForgotPassword",
		&data,
		true,
	)
	defer span.End()
	if err != nil {
		log.Error("[controllers.VerifyForgotPassword] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.VerifyForgotPassword(ctx, data.ForgotToken, data.Code)
	if err != nil {
		log.Error("[controllers.VerifyForgotPassword] cn.client.VerifyForgotPassword failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("data", response)
}

// ResetPassword reset password with token
func (cn *Controllers) ResetPassword(c *gin.Context) {
	log := hclog.Default()

	var data models.ResetPasswordRequest
	ctx, span, err := utils.HandleRequest(
		c,
		cn.tracer,
		"ResetPassword",
		&data,
		true,
	)
	defer span.End()
	if err != nil {
		log.Error("[controllers.ResetPasswordRequest] utils.HandleRequest failed", "error", err)
		return
	}

	response, err := cn.client.ResetPassword(ctx, data.ResetPasswordToken, data.NewPassword)
	if err != nil {
		log.Error("[controllers.ResetPassword] cn.client.ResetPassword failed", "error", err)
		c.Set("status", "failed")
		c.Set("code", status.Convert(err).Code())
		c.Set("message", status.Convert(err).Message())
		return
	}

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
	c.Set("data", response)
}
