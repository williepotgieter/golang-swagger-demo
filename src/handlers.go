package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHello() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	}
}

func GetAllProducts() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, GenerateProducts(100))
	}
}

func InitHandlers(r *gin.Engine) {
	r.GET("/", HandleHello())
	r.GET("/products", GetAllProducts())
}
