package main

import (
	"log"
	"net/http"

	"github.com/XiaoBeiAjie/aicongyou/config"
	"github.com/XiaoBeiAjie/aicongyou/db"
	"github.com/gin-gonic/gin"
)

func init() {
	err := config.Load("config/config.yaml")
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    err = db.Init()
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    defer db.Close()
}

func main() {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{
		"message": "ok",
		})
	})
	router.Run() // 默认监听 0.0.0.0:8080
}