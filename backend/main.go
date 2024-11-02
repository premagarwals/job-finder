package main

import (
	"time"

	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	rateLimiter := middleware.NewRateLimiter(10, time.Minute)
	r.Use(rateLimiter.LimitRequests)

	r.GET("/", controllers.Index)
	r.POST("/jobs", controllers.JobList)
	r.GET("/job/:id", controllers.JobRead)

	r.POST("/login", controllers.Login)

	auth := middleware.AuthMiddleware()
	r.POST("/job", auth, controllers.JobCreate)
	r.PUT("/job/:id", auth, controllers.JobUpdate)
	r.DELETE("/job/:id", auth, controllers.JobDelete)

	r.Run()
}
