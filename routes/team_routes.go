package routes

import (
	"jalcode-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupTeamRoutes(r *gin.Engine) {
	// t grup route dengan prefix /api/teams
	teamGroup := r.Group("/api/teams")
	{
		teamGroup.GET("/", controllers.GetTeamMembers)
		teamGroup.POST("/", controllers.CreateTeamMember)
		teamGroup.PUT("/:id", controllers.UpdateTeamMember)
		teamGroup.DELETE("/:id", controllers.DeleteTeamMember)
	}
}