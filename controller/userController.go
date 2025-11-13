package controller

import (
	"net/http"

	"github.com/XiaoBeiAjie/aicongyou/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) () {
	users, err := models.GetAllUsers()
      if err != nil {
          ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      }
      ctx.JSON(http.StatusOK, users)

}