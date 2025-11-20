package routes

import (
	"github.com/XiaoBeiAjie/aicongyou/controller"
	"github.com/gin-gonic/gin"
)

// SetupCauseRoutes 设置课程相关路由
func SetupCauseRoutes(r *gin.Engine) {
	causeGroup := r.Group("/api/causes")
	{
		// 获取所有课程
		causeGroup.GET("/", controller.GetAllCauses)

		// 根据ID获取课程
		causeGroup.GET("/:id", controller.GetCauseByID)

		// 根据参数查询课程（支持多条件和分页）
		causeGroup.GET("/search", controller.GetCausesByParams)

		// 可以在这里添加其他路由
		// causeGroup.POST("/", controller.CreateCause)
		// causeGroup.PUT("/:id", controller.UpdateCause)
		// causeGroup.DELETE("/:id", controller.DeleteCause)
	}
}