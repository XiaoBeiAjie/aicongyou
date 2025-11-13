package main

import (
	"log"
	"net/http"

	"github.com/XiaoBeiAjie/aicongyou/config"
	"github.com/XiaoBeiAjie/aicongyou/db"
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
	router.Run() // 默认监听 0.0.0.0:8080
    defer db.Close()
}