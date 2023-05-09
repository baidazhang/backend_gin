package routes

import (
	"backend_gin/controller"
	"backend_gin/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoverMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	return r
}
