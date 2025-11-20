package main

import (
	"log"
	"net/http"
	"os"

	"github.com/XiaoBeiAjie/aicongyou/config"
	"github.com/XiaoBeiAjie/aicongyou/db"
	"github.com/XiaoBeiAjie/aicongyou/routes"
	"github.com/gin-gonic/gin"
	"github.com/XiaoBeiAjie/aicongyou/controller"
)

func Init() {
	err := config.Load("config.yaml")
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    err = db.Init()
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
}

func main() {
	Init()
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{
		"message": "ok",
		})
	})
	router.GET("/user/getAll", func(ctx *gin.Context) {
		controller.GetAllUsers(ctx)
	})

	// 设置课程路由
	routes.SetupCauseRoutes(router)

	// 从环境变量获取端口，默认8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("服务器启动在端口: %s", port)
	router.Run(":" + port)
    defer db.Close()
}