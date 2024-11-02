package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/premagarwals/job-finder/initializers"
	"github.com/premagarwals/job-finder/models"
)

func JobDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Job{}, id)
	c.Status(200)
}
