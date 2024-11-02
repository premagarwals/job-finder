package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/premagarwals/job-finder/initializers"
	"github.com/premagarwals/job-finder/models"
)

func JobRead(c *gin.Context) {
	var job models.Job
	id := c.Param("id")
	initializers.DB.First(&job, id)

	c.JSON(200, gin.H{
		"job": job,
	})
}
