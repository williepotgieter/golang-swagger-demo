package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	InitHandlers(r)

	r.Run(":5000")
}
