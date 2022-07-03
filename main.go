package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	// NOTE: router に必要な設定を入れる
	router.POST("/login", func(ctx *gin.Context) {
		var json Login
		if err := ctx.ShouldBindJSON(&json); err != nil {
			// NOTE: err が詳細に鳴くと、攻撃者に仕組みを類推されるおそれがある
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// NOTE: router を起動する
	router.Run(":8080")
}
