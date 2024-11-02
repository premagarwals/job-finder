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
	r.GET("/job", controllers.JobList)
	r.GET("/job/:id", controllers.JobRead)

	r.POST("/login", controllers.Login)

	auth := middleware.AuthMiddleware()
	r.POST("/job", auth, controllers.JobCreate)
	r.PUT("/job/:id", auth, controllers.JobUpdate)
	r.DELETE("/job/:id", auth, controllers.JobDelete)

	r.Run()
}
