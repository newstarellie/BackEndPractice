// cd database_example
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

	// 創建表格
	createTableStmt := `
		CREATE TABLE IF NOT EXISTS example_table (
			column1 TEXT,
			column2 TEXT,
			column3 TEXT
		);
	`

	_, err = db.Exec(createTableStmt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("成功創建表格")

	// 插入資料
	insertStmt := `
		INSERT INTO example_table (column1, column2, column3)
		VALUES (?, ?, ?)
	`

	// 替換下面這一行中的 value1, value2, value3
	value1 := "your_value1"
	value2 := "your_value2"
	value3 := "your_value3"

	_, err = db.Exec(insertStmt, value1, value2, value3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("成功插入資料")
}
