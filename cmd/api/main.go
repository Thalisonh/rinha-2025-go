package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thalisonh/rinha-go/internal/adapter/handler"
)

func main() {
	r := gin.Default()

	handler.RegisterRoutes(r)

	r.Run(":8080")
} 