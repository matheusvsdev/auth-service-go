package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvsdev/auth-service-go/internal/repository"
	"github.com/matheusvsdev/auth-service-go/internal/transport/http"
)

func main() {
	db, err := repository.ConnectDB()
	if err != nil {
		panic("Erro ao conectar com banco: " + err.Error())
	}

	userRepo := &repository.UserRepository{DB: db}
	handler := &http.Handler{UserRepo: userRepo}

	r := gin.Default()
	r.POST("/register", handler.Register)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(":8080")
}
