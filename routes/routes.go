package routes

import (
	"go-test/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	pathAll := "/tasks"
	pathID := "/tasks/:id"
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET(pathAll, controllers.FindTasks)
	r.POST(pathAll, controllers.CreateTask)
	r.GET(pathID, controllers.FindTask)
	r.PATCH(pathID, controllers.UpdateTask)
	r.DELETE(pathID, controllers.DeleteTask)
	return r
}
