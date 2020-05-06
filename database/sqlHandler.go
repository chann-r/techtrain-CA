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
func NewSqlHandler() *SqlHandler {
  //接続処理
  conn, err := sql.Open("mysql", "root:rootpass@tcp(mysql:3306)/dojo_db")
  if err != nil {
    panic(err.Error)
  }

  sqlHandler := new(SqlHandler)
  sqlHandler.Conn = conn
  return sqlHandler
}

type SqlRow struct {
  Rows *sql.Rows
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (*SqlRow, error) {
  rows, err := handler.Conn.Query(statement, args...)
  if err != nil {
    return new(SqlRow), err
  }
  row := new(SqlRow)
  row.Rows = rows
  return row, nil
}

func (row SqlRow) Close() error {
  return row.Rows.Close()
}

func (row SqlRow) Next() bool {
  return row.Rows.Next()
}

func (row SqlRow) Scan(dest ...interface{}) error {
  return row.Rows.Scan(dest...)
}
