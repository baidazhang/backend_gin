package main

import (
	"backend_gin/controller"
	"backend_gin/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.AuthMiddleware(), middleware.RecoverMiddleware())
	r.POST("/api/auth/register", controller.Register)
	return r
}
