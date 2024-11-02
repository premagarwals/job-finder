package controllers

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.String(200, "Hello mate! Best wishes for your new job :)")
}
