// go run database_example.go

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 開啟 SQLite 資料庫（如果不存在會自動建立），連接資料庫
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // 程式結束前確保資料庫連線被正確關閉

	// 檢查是否可以連接到資料庫
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 連接成功，輸出成功訊息
	fmt.Println("成功連接到資料庫")
}
