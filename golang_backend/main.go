package main

import (
	"golang_backend/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Tambahkan route untuk endpoint AI
	r.POST("/ask", api.AskHandler)

	// Jalankan server pada port 8080
	r.Run(":8080")
}
