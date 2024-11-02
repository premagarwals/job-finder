package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/premagarwals/job-finder/controllers"
	"github.com/premagarwals/job-finder/initializers"
	"github.com/premagarwals/job-finder/middleware"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()

	rateLimiter := middleware.NewRateLimiter(10, time.Minute)
	r.Use(rateLimiter.LimitRequests)

	r.GET("/", controllers.Index)
	r.POST("/job", controllers.JobCreate)
	r.GET("/job", controllers.JobList)
	r.GET("/job/:id", controllers.JobRead)
	r.PUT("/job/:id", controllers.JobUpdate)
	r.DELETE("/job/:id", controllers.JobDelete)

	r.Run()
}
