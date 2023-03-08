package main

import (
	"github.com/gin-gonic/gin"
	"main.go/packs"
)

func main() {
	memoryStorage := packs.NewMemoryStorage()
	handler := packs.NewHandler(memoryStorage)

	router := gin.Default()

	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)

	router.Run()
}
