package database

import (
  // database/sqlパッケージ
  "database/sql"
  // mysql用のドライバー
  _ "github.com/go-sql-driver/mysql"
)

// DBと接続するための構造体
type SqlHandler struct {
  Conn *sql.DB
}

// DBと接続するための関数
func NewSqlHandler() SqlHandler {
  //接続処理
  conn, err := sql.Open("mysql", "root:rootpass@tcp(db:3306)/dojo_db")
  if err != nil {
    panic(err.Error)
  }

  sqlHandler := new(SqlHandler)
  sqlHandler.Conn = conn
  return sqlHandler
}
