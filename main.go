package main

import (
	"jalcode-api/config" 
	"jalcode-api/models"
	"jalcode-api/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// // Inisialisasi router Gin
	// r := gin.Default()

	// // Buat satu endpoint GET sederhana
	// r.GET("/api/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Halo! Server Go-mu sudah berjalan dengan lancar 🚀",
	// 	})
	// })

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.TeamMember{}, &models.Project{})

	r := gin.Default()

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server Go berjalan lancar 🚀"})
	})

	// route tim di sini
	routes.SetupTeamRoutes(r)

	//  server di port 8080
	r.Run(":8080")
}