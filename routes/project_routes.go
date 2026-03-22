package routes

import (
	"jalcode-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupProjectRoutes(r *gin.Engine) {
	projectGroup := r.Group("/api/projects")
	{
		projectGroup.GET("/", controllers.GetProjects)
		projectGroup.POST("/", controllers.CreateProject)
		projectGroup.PUT("/:id", controllers.UpdateProject)
		projectGroup.DELETE("/:id", controllers.DeleteProject)
	}
}

