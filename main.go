package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/data", func(c *gin.Context) {
		// 在這裡處理後端邏輯，並回傳 JSON 資料
		data := map[string]string{"message": "Hello from Go!"}
		c.JSON(http.StatusOK, data)
	})

	router.Run(":8080") // 這會啟動伺服器在本地的 8080 port
}
