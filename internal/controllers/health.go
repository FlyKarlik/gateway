package controllers

import (
	"comet/utils"
	"github.com/gin-gonic/gin"
)

// Status state of service
func (cn *Controllers) Status(c *gin.Context) {
	tr := cn.tracer
	_, span := tr.Start(c, "Health")
	defer span.End()

	c.Set("status", "success")
	c.Set("code", utils.CodeOK)
}
