// cd API
// go run getData.go

// 程式碼用途 : 使用 Go 語言和 Gin 框架建立的簡單的 HTTP API，它提供了一個端點 /api/data 來擷取資料庫中某個表格

// TODO : 目前只會出現 {"error":"資料庫查詢失敗"}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 開啟 SQLite 資料庫（如果不存在會自動建立），連接資料庫
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 檢查是否可以連接到資料庫
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("成功連接到資料庫")

	// 使用 Gin 框架建立 HTTP 路由
	router := gin.Default()

	// 定義一個 API 路由來取得資料
	router.GET("/api/data", func(c *gin.Context) {
		// 執行資料庫查詢
		rows, err := db.Query("SELECT column1, column2, column3 FROM example_table")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "資料庫查詢失敗"})
			return
		}
		defer rows.Close()

		// 準備一個 slice 來存放查詢結果
		var data []map[string]string

		// 迭代查詢結果，將每一行的值加入到 slice 中
		for rows.Next() {
			var column1, column2, column3 string
			err := rows.Scan(&column1, &column2, &column3)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "讀取資料失敗"})
				return
			}

			rowData := map[string]string{
				"column1": column1,
				"column2": column2,
				"column3": column3,
			}
			data = append(data, rowData)
		}
		
		// 檢查是否有資料
		if len(data) == 0 {
				c.JSON(http.StatusOK, gin.H{"message": "資料庫是空的"})
				return
		}

		// 將查詢結果返回給客戶端
		c.JSON(http.StatusOK, gin.H{"data": data})
	})

	// 啟動 HTTP 伺服器
	router.Run(":8080")
}
