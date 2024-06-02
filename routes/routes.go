package routes

import (
	"github.com/DLLenjoyer/AdminApp-with-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", func(c *gin.Context) {
			controllers.GetAllUsers(c, db)
		})
	}
}
