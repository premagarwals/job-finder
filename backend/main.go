package main

import (
	"github.com/gin-gonic/gin"
	"github.com/premagarwals/job-finder/controllers"
	"github.com/premagarwals/job-finder/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()

	r.GET("/", controllers.Index)
	r.POST("/job", controllers.JobCreate)
	r.GET("/job/:id", controllers.JobRead)

	r.Run()
}
